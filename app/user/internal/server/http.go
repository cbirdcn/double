package server

import (
	v1 "double/api/user/v1"
	"double/app/user/internal/biz"
	"double/app/user/internal/conf"
	"double/app/user/internal/service"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	httpNet "net/http"
	"time"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, user *service.UserService, logger log.Logger, useCase *biz.UserUseCase) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			CheckToken(useCase),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// 强制改变返回结构，目的是避免框架将httpCode和业务Code绑定在一起。反restful结构，让httpCode固定=200，业务Code自定义（正常默认200），更适合无限扩展的错误类型。
	// 正常ResponseWriter结构举例：{"Code":200,"Data":{"accountDetail":{"id":"62dadb272d26ce27565c318f","accountName":"Hello","createdAt":"2022-07-22 17:15:19","status":1}},"Ts":"2022-07-22 17:15:19.06206721 +0000 UTC"}
	opts = append(opts, http.ResponseEncoder(func(
		w httpNet.ResponseWriter,
		r *httpNet.Request,
		i interface{},
	) error {
		type response struct {
			Code int
			MetaData interface{}
			Data interface{}
			Ts   string
		}
		reply := &response{
			Code: 200,
			MetaData: w.Header(), // TODO: test
			Data: i,
			Ts:   time.Now().Local().String(),
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(reply)
		if err != nil {
			return err
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
		return nil
	}))

	// 错误结构举例：{"Code":404,"Reason":"USER_NAME_EXIST","Message":"account exist: Hello","Ts":"2022-07-22 17:14:05.608411034 +0000 UTC"}
	opts = append(opts, http.ErrorEncoder(func(
		w httpNet.ResponseWriter,
		r *httpNet.Request,
		e error,
	) {
		type response struct {
			Code int32
			MetaData interface{}
			Reason   string
			Message   string
			Ts   string
		}
		reply := &response{
			Code: e.(*errors.Error).GetCode(),
			MetaData: w.Header(), // TODO: test
			Reason: e.(*errors.Error).GetReason(),
			Message: e.(*errors.Error).GetMessage(),
			Ts:   time.Now().Local().String(),
		}
		codec := encoding.GetCodec("json")
		data, err := codec.Marshal(reply)
		log := log.NewHelper(logger)
		if err != nil {
			log.Errorf("failed ping connection pool to mysql: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write(data)
	}))

	srv := http.NewServer(opts...)

	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	v1.RegisterUserHTTPServer(srv, user)
	return srv
}

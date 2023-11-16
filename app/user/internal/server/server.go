package server

import (
	"context"
	user_v1 "double/api/user/v1"
	"double/app/user/internal/biz"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/wire"
	"strconv"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer)

const IGNORE_OPERATION_LIST = ["/api.user.v1.User/CreateAccount"]

// 自定义中间件
func CheckToken(useCase *biz.UserUseCase) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				fmt.Println(tr.Operation())
				op := tr.Operation()
				if ()
				accountId := tr.RequestHeader().Get("x-account-id")
				token := tr.RequestHeader().Get("x-token")
				accountIdInt, _ := strconv.ParseInt(accountId, 10, 64)
				user, _ := useCase.Get(ctx, accountIdInt)
				// 检查token正确性
				if user.UserName != token {
					return nil, user_v1.ErrorUserPasswordDifferent("account id: %s", accountId)
				}
				// 检查封号
				if user.IsBanned {
					return nil, user_v1.ErrorUserBanned("account id: %s", accountId)
				}

				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
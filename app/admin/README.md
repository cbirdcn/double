# Admin Service

后台服务

需要调用其他service，比如调用user service的ListUser()

不再使用本地配置，使用配置中心，需要转码

通过jaeger查看服务间trace过程

# User Service

第一个service

管理用户

## 调用user service请求响应举例

## ListUser

url
```shell
curl 127.0.0.1:8002/admin/v1/user
```
req
```json
{}
```
res
```json
{
  "userDetails": [
    {
      "id": "1",
      "userName": "user1",
      "clubName": "club1",
      "createdAt": "2021-12-15 02:47:52"
    },
    {
      "id": "3",
      "userName": "user3",
      "clubName": "club3",
      "createdAt": "2021-12-15 02:47:52"
    },
    {
      "id": "2",
      "userName": "user2",
      "clubName": "club2",
      "createdAt": "2021-12-15 02:47:52"
    },
    ...
  ]
}
```
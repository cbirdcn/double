# service请求响应举例

## CreateUser
req
```json
{
  "userName": "Hello",
  "password": "Hello",
  "clubName": "Hello"
}
```
res
```json
{
    "userDetail": {
        "id": "1",
        "userName": "Hello",
        "clubName": "Hello",
        "createdAt": "2021-12-14 09:48:10"
    }
}
```

## BanUser

```json
{
  "id": 1
}
```

```json
{
  "res": true
}
```

## UpdateUserName

```json
{
  "id": 1,
  "newUserName": "Hello2"
}
```

```json
{
  "res": true
}
```

## UpdateUserName

```json
{
  "id": 1,
  "oldUserPassword": "Hello",
  "newUserPassword": "Hello2"
}
```

```json
{
  "res": true
}
```

## UpdateUserPassword

```json
{
  "id": 1,
  "oldUserPassword": "Hello2",
  "newUserPassword": "Hello3"
}
```

```json
{
  "res": true
}
```

## GetUser

```json
{
  "id": 1
}
```

```json
{
  "userDetail": {
    "id": "1",
    "userName": "Hello2",
    "clubName": "Hello",
    "createdAt": "2021-12-14 09:48:10"
  }
}
```

## UserLogin

```json
{
  "userName": "Hello2",
  "password": "Hello3"
}
```

```json
{
  "res": true,
  "userDetail": {
    "id": "1",
    "userName": "Hello2",
    "clubName": "Hello",
    "createdAt": "2021-12-14 09:48:10"
  }
}
```

## ListUser

```json
{}
```

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
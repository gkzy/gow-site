#####  支持的HTTPMethod
```go
HTTPMethod = map[string]bool{
    "GET":     true,
    "POST":    true,
    "PUT":     true,
    "DELETE":  true,
    "PATCH":   true,
    "OPTIONS": true,
    "HEAD":    true,
    "TRACE":   true,
}
```

##### 使用方式

包括基本路由与分组

```go
r := gow.Default()

r.GET(path,handler)
r.POST(path,handler)
r.PUT(path,handler)
r.DELETE(path,handler)
r.PATCH(path,handler)
r.OPTIONS(path,handler)
r.HEAD(path,handler)
r.TRACE(path,handler)
```

##### 路由参数

```go
r.Any("/article/:id", ArticleDetailHandler)
```

*获取 param 值*

```go
id:=c.Param("id")
```


##### 路由分组


*main.go*

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    v1 := r.Group("/v1")
    {
        v1.GET("/user/:id", GetUser)
        v1.DELETE("/user/:id", DeleteUser)
    }

    r.Run()
}

func GetUser(c *gow.Context) {
    c.JSON(gow.H{
        "nickname": "新月却泽滨",
        "qq":       "301109640",
    })
}

func DeleteUser(c *gow.Context) {
    c.JSON(gow.H{
        "code": 0,
        "msg": "success",
    })
}

```

*Get*

```shell
curl -i http://127.0.0.1:8080/v1/user/1


HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Thu, 16 Jul 2020 05:55:16 GMT
Content-Length: 46

{
  "nickname": "新月却泽滨",
  "qq": "301109640"
}

```

*Delete*

```shell
curl  -X "DELETE" http://127.0.0.1:8080/v1/user/1

{
    "code":0,
    "msg":"success"
}

```
## 路由设置

###  支持的请求方式
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

### 基础路由


```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()

    r.GET("/someGet", getting)
    r.POST("/somePost", posting)
    r.PUT("/somePut", putting)
    r.DELETE("/someDelete", deleting)
    r.PATCH("/somePatch", patching)
    r.HEAD("/someHead", head)
    r.OPTIONS("/someOptions", options)
    r.Run()
}
```

*典型例子*

```go
package main

import (
    "github.com/gkzy/gow"
)

func main(){
    r.GET("/", func(c *gow.Context) {
        c.JSON(gow.H{
            "code": 0,
            "msg":  "success",
        })
    })
}
```


#### 路由参数


```go
r.GET("/article/:id", handler)
```

*获取 param 值*

```go
id:=c.Param("id")
```


### 分组路由


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


    v2 := r.Group("/v2")
    {
        v2.GET("/user/:id", GetUser)
        v2.DELETE("/user/:id", DeleteUser)
    }

    r.Run()
}

func GetUser(c *gow.Context) {
 
}

func DeleteUser(c *gow.Context) {

}

```

支持middleware中间件模式，

```go
r := gow.Default()
r.Use(...)
```

##### 内置默认使用的
```go
func Default() *Engine {
    engine := New()
    engine.Use(Recovery())  //Recovery
    engine.Use(Logger())   //Logger
    return engine
}

```

##### 比如需要一个鉴权的中间件

```go
package main

import "github.com/gkzy/gow"


func main(){
    r := gow.Default()
    v1:=r.Group("/v1")          //v1的路由
    //v1.Use(APIAuth())         //实现对/v1/* api的鉴权
    user:=v1.Group("/user")     //user的路由
    user.Use(APIAuth())         //只实现对/v1/user/* api的鉴权
    //user.User(UserAuth())     //另外一个中间件调用 
    r.Run()
}


// APIAuth 简单的鉴权方法
func APIAuth() gow.HandlerFunc {
    return func(c *gow.Context) {
        //取header传递的token
        token := c.GetHeader("token") 
        //验证token
        if token != "123456" {  
            c.ServerJSON(403, gow.H{
                "code": 403,
                "msg":  "没有访问权限",
            })
            return
        }
        //需要执行此方法，否则不会执行后面的请求
        c.Next() 
    }
}

// UserAuth 用户鉴权
func UserAuth() gow.HandlerFunc {
    return func(c *gow.Context) {
        ......
        .....
        //需要执行此方法，否则不会执行后面的请求
        c.Next() 
    }
}

```
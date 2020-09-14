### 在gow中使用session



#### 创建演示代码

```shell
mkdir demo_session
cd demo_session
go mod init

vi main.go
```


**main.go**

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    gow.InitSession()  
    r.Use(gow.Session())
    r.GET("/", SetUser)
    r.GET("/get", GetUser)
    r.GET("/del", DeleteUser)
    r.Run() //:8080
}

var (
    key = "user-token"
)

func SetUser(c *gow.Context) {
    c.SetSession(key, "session_value")  //设置
}

func GetUser(c *gow.Context) {
    v := c.GetSessionString(key)    //获取
    c.String(v)
}

func DeleteUser(c *gow.Context) {
    c.DeleteSession(key)        //删除
}
```

**执行**
```shell
go run main.go
```

**分别访问**

```go
http://127.0.0.1:8080       // set session
http://127.0.0.1:8080/get   // get session
http://127.0.0.1:8080/del   // delete session
```

#### 初始化session模块

```go
gow.InitSession()
```

##### 注册中间件

```go
r.Use(gow.Session())
```
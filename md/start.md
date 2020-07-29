**gow** 是一个基于gin框架思想和beego框架中html模板处理机制，封装的一个go web框架。可用于开发Web API和Web网站项目。


#### 项目地址：

[https://github.com/gkzy/gow](https://github.com/gkzy/gow)

#### 官网地址：

[https://gow.22v.net](https://gow.22v.net)


---

###  快速开始

```shell
mkdir hello
cd hello
```

```shell
go mod init
```

```shell
go get github.com/gkzy/gow
```

##### 创建 main.go

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()

    r.GET("/", func(c *gow.Context) {
        c.JSON(gow.H{
            "code": 0,
            "msg":  "success",
        })
    })

    //default :8080
    r.Run()

    //r.Run(":9090") 
    //r.Run("127.0.0.1:9090")
}
```

#####  编译和运行
```shell
go build && ./hello
或
go run main.go
```

#####  访问地址

```shell
浏览器访问：http://127.0.0.1:8080
```

```shell
curl -i http://127.0.0.1:8080
```

```shell
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Wed, 15 Jul 2020 09:15:31 GMT
Content-Length: 27

{   
    "code":0,
    "msg":"success"
}
```

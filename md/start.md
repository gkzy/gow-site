
# gow

gow 是基于gin的HTTP框架，在gin的基础上，做了更适合的html模板封装和数据输出。可用于开发Web API和Web网站项目

## 项目地址：

[https://github.com/gkzy/gow](https://github.com/gkzy/gow)

#### 官网地址：

[https://gow.22v.net](https://gow.22v.net)


---

##  快速开始

### 安装golang

```sh
https://golang.org/dl/
```

### 安装gow

```sh
go get github.com/gkzy/gow
```

### 开始一个新项目

```sh
mkdir hello
cd hello
```

```shell
go mod init
```



#### 创建 main.go

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

    //r.Run(9090)
    //r.Run(":9090") 
    //r.Run("127.0.0.1:9090")
}
```

####  编译和运行
```shell
go run main.go
或
go build && ./hello
```

####  访问地址

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

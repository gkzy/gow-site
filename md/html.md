
## 做个网站

包括HTML模板文件处理和静态资源处理

### 目录结构

```
PROJECT_NAME
├──static
      ├── img
            ├──111.jpg
            ├──222.jpg
            ├──333.jpg
      ├──js
      ├──css
├──views
    ├──index.html
    ├──article
        ├──detail.html
├──main.go
```

###  设置模板目录

```go
r := gow.Default()
r.SetView("views")
```

### 打开模板渲染

```go
r.AutoRender = true
```

###  设置静态资源

```go
r.Static("/static", "static")
```

### 设置favicon.ico

```go
r.StaticFile("favicon.ico","static/img/log.png") 
```

### 添加自定义模板函数
```go
r.AddFuncMap(key string,fn interface(){})
```

### 设置html模板符号 

```go
r.SetDelims("{{","}}")
```

### 向模板输出数据方法一：

```go
c.HTML("article/detail.html", gow.H{
    "title":    "年薪百万的文科专业有哪些？",
})
```

``` css
<title>{{.title}}</title>
```

### 向模板输出数据方法二：

```go
c.Data["title"] = "年薪百万的文科专业有哪些？"
c.HTML("article/defail.html")
```

``` css
<title>{{.title}}</title>
```

###  演示代码

*main.go*

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    r.AutoRender = true //打开html模板渲染
    r.SetView("views") //默认静态目录为views时，可不调用此方法
    r.StaticFile("favicon.ico","static/img/log.png")  //路由favicon.ico
    r.Static("/static", "static")

    //router
    r.Any("/", IndexHandler)
    r.Any("/article/1", ArticleDetailHandler)


    //自定义hello的模板函数
    //在模板文件可通过 {{hello .string}} 来执行
    r.AddFuncMap("hello", func(str string) string {
        return "hello:" + str
    })

    r.Run()
}

//IndexHandler 首页
func IndexHandler(c *gow.Context) {
    c.HTML("index.html", gow.H{
        "name":    "gow",
        "package": "github.com/gkzy/gow",
    })
}

//ArticleDetailHandler 文件详情页
func ArticleDetailHandler (c *gow.Context){
    c.HTML("article/detail.html", gow.H{
        "title":    "年薪百万的文科专业有哪些？",
    })
}
```



#### 运行

```shell
go run main.go
```

#### 浏览器访问

```shell
https://127.0.0.1:8080/
https://127.0.0.1:8080/article/1
```
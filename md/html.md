包括HTML模板文件处理和静态资源处理


##### 目录结构

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

#####  设置模板目录

```go
r.SetView("views")
```

#####  设置静态资源

```go
r.Static("/static", "static")
```

#####  演示代码

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

//IndexHandler 首页handler
func IndexHandler(c *gow.Context) {
    c.HTML("index.html", gow.H{
        "name":    "gow",
        "package": "github.com/gkzy/gow",
    })
}

//ArticleDetailHandler 文章详情页handler
func ArticleDetailHandler (c *gow.Context){
    c.HTML("article/detail.html", gow.H{
        "title":    "年薪百万的文科专业有哪些？",
    })
}
```

*views/index.html*

```html
<html>
<head>
    <title>{{hello .name}}</title>
    <meta charset="utf-8"/>
</head>
<body>
    <h2>{{.name}}</h2>
    <hr>
    <h5>{{.package}}</h5>
</body>
</html>
```

*views/article/detail.html*

```html
<html>
<head>
    <title>{{.title}}</title>
    <meta charset="utf-8"/>
    <style>
        img{max-width:600px;}
    </style>
</head>
<body>
    <h2>{{.title}}</h2>
    <hr>
    <p><img src="/static/img/111.jpg"></p>
    <p><img src="/static/img/222.jpg"></p>
    <p><img src="/static/img/333.jpg"></p>
</body>
</html>
```

##### 运行

```shell
go run main.go
或
go build main.go -o app && ./app
```

##### 浏览器访问

```shell
https://127.0.0.1:8080/
https://127.0.0.1:8080/article/1
```
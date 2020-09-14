
## 参数配置

可以通过两种式来完成项目的基础配置

### 1. 使用代码

#### 运行模式
```go
r := gow.Default()
r.RunMode = "dev"
```

#### 打开模板渲染

```go
r.AutoRender = true
```

#### 设置模板目录

```go
r.SetView("views")
```

#### 静态资源
> 图片、字体、css、js等

```go
r.Static("/static", "static")
```
#### 配置html函数调用时的符号
```go
r.SetDelims("{{","}}")  
```

#### 添加自定义模板Func

```go
r.AddFuncMap(key string, fn interface{})
```
#### session 开关

```go
r.SetSessionOn(true)
```

## 2. 配置文件


### 目录结构

```shell
PROJECT_NAME
├──conf
    ├──app.conf # 默认的配置文件 
    ├──prod.app.conf # 环境变量 APP_RUN_MODE="prod" 时的配置文件
    ├──dev.app.conf  # 环境变量 APP_RUN_MODE="dev" 时的配置文件
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

### 使用gow的统一配置

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    // gow.GetAppConfig 会读取当前所使用配置文件中的基础配置
    // 再通过r.SetAppConfig设置
    r.SetAppConfig(gow.GetAppConfig())  
    r.Run()
}
```


#### 配置文件

*app.conf*

> 开发环境

```shell
app_name = gow-site
run_mode = dev
http_addr = 8080
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

> 生产环境

```shell
app_name = gow-site
run_mode = prod  # 此处改为prod
http_addr = 8080
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

#### 使用环境变量时

*prod.app.conf*

> 仅在使用 GOW_RUN_MODE="prod" 时生效

``` shell
pp_name = gow-site
run_mode = prod
http_addr = 8080
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

*dev.app.conf*

> 仅在使用 GOW_RUN_MODE="dev" 时生效

``` shell
pp_name = gow-site
run_mode = dev
http_addr = 8080
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```



---


### 实现自定义配置

使用的库:

```sh
github.com/gkzy/gow/lib/config
```

*app.conf*

```shell

api_host = "https:/api.22v.net"

[redis]
host = "192.168.0.197"
port = 6379
db = 0
password = "123456"
maxidle = 50
maxactive = 10000
```

*redis配置信息读取*

```go
package main

import "github.com/gkzy/gow/lib/config"

func main(){
    GetConfig()
}


func GetConfig(){
    apiHost:=config.GetString("api_host")
    // 获取带Section的参数
    host:=config.DefaultString("redis::host","192.168.0.197")
    port:=config.DefaultInt("redis::port",6379)
    db:=config.DefaultString("redis::db",0)
    password:=config.DefaultString("redis::password","123456")
    maxidle:=config.DefaultInt("redis::maxidle",50)
    maxactive:=config.DefaultInt("redis::maxactive",10000)
    fmt.Println(apiHost,host,prot,db,password,maxidle,maxactive)
}
```

*更多方法*
```go

//不带默认值
config.GetString()
config.GetInt()
config.GetInt64()
config.GetFloat()
config.GetBool()

//带默认值
config.DefaultString()
config.DefaultInt()
config.DefaultInt64()
config.DefaultFloat()
config.DefaultBool()

```
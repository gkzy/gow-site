可以在conf目录中，指定配置文件

使用到的库：

```shell
github.com/gkzy/gow/lib/config
```

使用到的第三方ini配置库：

```shell
github.com/go-ini/ini
```

#### 有一项目：

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


#### 仅使用app.conf时

*app.conf*

> 开发环境

```shell
app_name = gow-site
run_mode = dev
http_addr = ":8080"
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

> 生成环境

```shell
app_name = gow-site
run_mode = prod  # 此处改为prod
http_addr = ":8080"
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

#### 使用环境变量时

*prod.app.conf*

> 仅在使用 APP_RUN_MODE="prod" 时生效

``` shell
pp_name = gow-site
run_mode = prod
http_addr = ":8080"
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

*dev.app.conf*

> 仅在使用 APP_RUN_MODE="dev" 时生效

``` shell
pp_name = gow-site
run_mode = dev
http_addr = ":8080"
views = "views"
auto_render = true
recover_panic = true
template_left = "{{"
template_right = "}}"
```

#### 使用gow的统一配置

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

---

#### 实现其他配置信息读取

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

```go
package main

import "github.com/gkzy/gow/lib/config"


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

#####  更多方法
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
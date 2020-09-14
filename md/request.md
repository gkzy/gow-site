## 请求数据

### 获取参数

```go

func GetUser(c *gow.Context){

    //获取字串
    c.GetString("key","default")

    //获取int
    c.GetInt("key",0)

    //获取bool
    c.GetBool("key",false)

    //获取int64
    c.GetInt64("key",-1)

    //获取float
    c.GetFloat("key",0)

    //获取[]string
    var ids []string
    ids = c.GetStrings("ids")  

    //其他方法
    c.GetInt32()
    c.GetInt16()
    c.GetInt8()
    ....
}

```


### 获取request body的内容

```go

type User struct {
    Nickname string `json:"nickname"`
    QQ       string `json:"qq"`
}

func GetUser(c *Context){
    user := new(User)
    err := c.DecodeJSONBody(&user)
    if err != nil {
        //handler error
    }

    c.JSON(gow.H{
        "user": user,
    })   

}
```


### 文件上传

当需要上传大于32MB的文件时，请使用以下配置：

```sh
r.MaxMultipartMemory = 1<<22 //64MB
```

gow提供了三个方法来处理文件上传：  

```sh
1. GetFile(key string) (multipart.File, *multipart.FileHeader, error)  # 单个文件上传
2. GetFiles(key string) ([]*multipart.FileHeader, error)   #多个文件上传
3. SaveToFile(fromFile, toFile string) error     # //保存到服务器
```

``` css 
<form enctype="multipart/form-data" method="post">
    <input type="file" name="uploadname" />
    <input type="submit">
</form>
```

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    // gow.GetAppConfig 会读取当前所使用配置文件中的基础配置
    // 再通过r.SetAppConfig设置
    r.POST("/upload",UploadFile)
    r.Run()
}

func UploadFile(c *gow.Context){
    f,h,err:=c.GetFile("file")
    if err!=nil{
        log.Fatal("getfile err ", err)
    }
    defer f.Close()
    c.SaveToFile("file","upload/"+h.Filename) //保存在upload下，没有目录，需要先创建
}
```

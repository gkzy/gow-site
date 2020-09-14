
## 数据输出

支持：json jsonp xml yaml string 等方式输出

### struct

```go
type User struct{
    Nickname string `json:"nickname" xml:"nickname"`
    Age int `json:"age" xml:"ag"`
}
```

### JSON


```go
func GetUser(c *gow.Context){
    //default http.StatusOK
    c.JSON(gow.H{
        "nickname":"gow",
        "age":1,
    })

   //或者，指定 http.StatusCode
    c.ServerJSON(200,gow.H{
        "nickname":"gow",
        "age":1,
    })
}

```

### JSONP


```go
func GetUser(c *gow.Context){
    //default http.StatusOK
    c.JSONP(gow.H{
        "nickname":"gow",
        "age":1,
    })

   //或者，指定 http.StatusCode
    c.ServerJSONP(200,gow.H{
        "nickname":"gow",
        "age":1,
    })
}

```

### XML

```go
func GetUser(c *gow.Context){

    //default http.StatusOK
    c.XML(gow.H{
        "nickname":"gow",
        "age":18,
    })

   //或者，指定 http.StatusCode
    c.ServerXML(200,gow.H{
        "nickname":"gow",
        "age":18,
    })
}

```


### YAML

```go
func GetUser(c *gow.Context){

    //default http.StatusOK
    c.YAML(gow.H{
        "nickname":"gow",
        "age":18,
    })

   //或者，指定 http.StatusCode
    c.ServerYAML(200,gow.H{
        "nickname":"gow",
        "age":18,
    })
}

```

### String

```go
func GetUser(c *gow.Context){

    //default http.StatusOK
    c.String("hello gow...")

   //或者，指定 http.StatusCode
    c.ServerString(200,"hello gow...")
}

```

### File

```go
func GetUser(c *gow.Context){
    //读取并输出
    c.File("go.mod")
}
```

### Download

*下载指定内容*

```go

func GetUser(c *gow.Context){
    c.Download([]byte("download string"))
}
```

*下载文件文件*

```go
func GetUser(c *gow.Context){
    c.FileAttachment("main.go","main.txt")
}
```
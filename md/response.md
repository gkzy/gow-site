支持：json xml string html 等方式输出

**JSON**

```go
func GetUser(c *Context){

    //default http.StatusOK
    c.JSON(gow.H{
        "nickname":"gow",
        "age":18,
    })

   //或者，指定 http.StatusCode
    c.ServerJSON(200,gow.H{
        "nickname":"gow",
        "age":18,
    })
}

```

**XML**

```go
func GetUser(c *Context){

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

**String**

```go
func GetUser(c *Context){

    //default http.StatusOK
    c.String("hello gow...")

   //或者，指定 http.StatusCode
    c.ServerString(200,"hello gow...")
}

```

**File**


```go
func GetUser(c *Context){

    //读取并输出
    c.File("go.mod")
}
```
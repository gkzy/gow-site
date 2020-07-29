**请参考以下演示代码**

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


**获取 request body，并反序列化到 struct**

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
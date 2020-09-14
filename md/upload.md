#### 上传文件

使用gow上传文件



##### 上传到服务器

```shell
使用 gow.Context的SaveToFile()方法
```

```go
package main

import (
    "github.com/gkzy/gow"
)

func main() {
    r := gow.Default()
    r.POST("/upload", Upload)
    r.Run()
}

func Upload(c *gow.Context) {
    err := c.SaveToFile("file", "./upload/1.jpg")
    if err != nil {
        c.JSON(gow.H{
            "code": 1,
            "msg":  err.Error(),
        })
        return
    }

    c.JSON(gow.H{
        "code": 0,
        "msg":  "success",
    })
}

```

```shell
go run main
```

```shell
POST http://127.0.0.1:8080/upload
```

##### 上传到阿里云的对象存储

```shell
使用 gow.Context的GetFile方法
```

```go
package main

import (
    "fmt"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
    "github.com/gkzy/gow"
    "io"
)

func main() {
    r := gow.Default()
    r.POST("/upload", Upload)
    r.Run()
}

func Upload(c *gow.Context) {
    f, h, err := c.GetFile("file")

    if err != nil {
        c.JSON(gow.H{
            "code": 1,
            "msg":  err.Error(),
        })
        return
    }

    url, err := PutObject(f, "user", h.Filename)
    if err != nil {
        c.JSON(gow.H{
            "code": 1,
            "msg":  err.Error(),
        })
        return
    }

    c.JSON(gow.H{
        "code": 0,
        "msg":  "success",
        "url":  url,
    })
}

var (
    // 阿里里的一些上传配置，请自己完成
    //TODO:
    endPoint    string
    accessKeyId string
    accessKey   string
    bucketName  string
    imgServer   string
)

// PutObject 上传到aliyun oss
func PutObject(reader io.Reader, dir string, fileName string) ( string, error) {
    client, err := oss.New(endPoint, accessKeyId, accessKey)
    if err != nil {
        return "", err
    }
    bucket, err := client.Bucket(bucketName)
    if err != nil {
        return "", err
    }

    imgUrl := fmt.Sprintf("%s/%s", dir, fileName)
    err = bucket.PutObject(imgUrl, reader)
    if err != nil {
        return "", err
    }
    return imgServer + imgUrl, nil
}


```
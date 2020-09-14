#####  HTML模板函数

* str2html

```js
{{str2html .title}}
```

* html2str

```js
{{html2str .title}}
```

* substr

```js
{{substr .title 0 30}}
```

* datetimeformat

> now为time.Time

```js
{{datetimeformat .now "YYYY-MM-DD HH:mm:ss"}}
```

* date

> now为time.Time

```js
{{date .now}}
```

* int_datetimeformat

> int_now为int时间戳

```js
{{date .int_now "YYYY-MM-DD HH:mm:ss"}}
```

* int_datetime

> int_now为int时间戳

```js
{{int_datetime .int_now}}
```

* int_date

> int_now为int时间戳


```js
{{in_date .int_now}}
```

* assets_js

```js
{{assets_js "/static/js/main.js"}}
```

* assets_css

```js
{{assets_js "/static/css/gow.css"}}
```


##### 添加自己的模板函数


*定义函数*

```go
// hello 定义一个func
func hello(str string) string{
    return "hello:"+str
}
```

*添加函数*

```go
r := gow.Defualt()
r.AddFuncMap("hello",hello)
```
*在模板中使用*

```go
{{hello "Sam"}}

//会在html模板上，输出：hello:Sam
```
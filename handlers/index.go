package handlers

import (
	"fmt"
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-site/md"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"strings"
	"time"
)

var (
	version = time.Now().Unix()
	icons   = []string{"fe-flag", "fe fe-at-sign", "fe fe-globe", "fe-terminal", "fe-feather", "fe-image", "fe-pie-chart", "fe-check-square", "fe-tag", "fe-book-open"}
)

//IndexHandler 首页
func IndexHandler(c *gow.Context) {
	c.Redirect(302, "/docs/start")
	//c.HTML("index.html")
}

//DocsHandler 读文件并渲染
func DocsHandler(c *gow.Context) {

	data := gow.H{}
	url := c.Request.URL.String()
	mdFile := md.GetMDFile(url)
	b, err := ioutil.ReadFile(mdFile.FilePath)
	if err != nil {
		c.Status(404)
		return
	}
	mb := blackfriday.MarkdownCommon(b)
	data["content"] = string(mb)
	data["title"] = mdFile.Title
	data["menu"] = builderMenus(url)
	data["version"] = version
	c.HTML("docs.html", data)
}

//builderMenus 右侧菜单
func builderMenus(url string) string {
	data := md.GetMdFiles()
	var b strings.Builder
	for index, item := range data {
		var style string
		if strings.ToLower(url) == strings.ToLower(item.URL) {
			style = "active"
		}
		b.WriteString(fmt.Sprintf("<a href=\"%s\" class=\"list-group-item list-group-item-action %s\"><span class=\"icon mr-3\"><i class=\"fe %s\"></i></span>%s</a>", item.URL, style, icons[index], item.Title))
	}

	return b.String()
}

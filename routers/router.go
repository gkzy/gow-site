package routers

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-site/handlers"
	"github.com/gkzy/gow-site/md"
)

// HTMLRouter
func HTMLRouter(r *gow.Engine) *gow.Engine {
	//静态文件
	r.Static("/static","static")
	r.StaticFile("favicon.ico,","./static/img/icon.png")

	//错误处理
	r.NoRoute(handlers.NotFoundHandler)

	r.Any("/", handlers.IndexHandler)

	data := md.GetMdFiles()

	for _, item := range data {
		r.Any(item.URL, handlers.DocsHandler)
	}

	return r
}

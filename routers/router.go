package routers

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-site/handlers"
	"github.com/gkzy/gow-site/md"
)

// HTMLRouter
func HTMLRouter(r *gow.Engine)  {
	r.GET("/", handlers.IndexHandler)
	data := md.GetMdFiles()
	for _, item := range data {
		r.GET(item.URL, handlers.DocsHandler)
	}

}

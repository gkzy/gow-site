package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-site/routers"
)

func main() {
	r := gow.Default()
	r.SetAppConfig(gow.GetAppConfig())
	r = routers.HTMLRouter(r)
	r.Run()
}

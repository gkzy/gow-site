package main

import (
	"github.com/gkzy/gow"
	"github.com/gkzy/gow-site/routers"
	"github.com/gkzy/gow/lib/config"
	"github.com/gkzy/gow/lib/logy"
	"os"
)

const (
	ServerName = "gow-site"
)

func init() {

	gow.InitConfig()

	InitLog()

}

func main() {
	r := gow.Default()
	r.Static("/static", "./static")
	r.SetAppConfig(gow.GetAppConfig())
	routers.HTMLRouter(r)
	r.Run()
}

// InitLog init logy
func InitLog() {
	runMode := config.DefaultString("run_mode", "dev")
	if runMode == gow.ProdMode {
		logy.SetOutput(
			logy.MultiWriter(
				logy.WithColor(logy.NewWriter(os.Stdout)),
				logy.NewFileWriter(logy.FileWriterOptions{
					Dir:           "./logs",
					Prefix:        "web",
					StorageMaxDay: 7,
					StorageType:   logy.StorageTypeDay,
				}),
			),
			ServerName,
		)
	} else {
		logy.SetOutput(
			logy.WithColor(logy.NewWriter(os.Stdout)),
			ServerName,
		)
	}

	logy.Infof("-------------------------------------------")
	logy.Infof("Start %s ......", ServerName)
	logy.Infof("-------------------------------------------")
}

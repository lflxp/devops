package main

import (
	"github.com/lflxp/boltDb/pkg/db"
	_ "github.com/lflxp/boltDb/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		// ui
		beego.BConfig.WebConfig.StaticDir["/"] = "ui/dist"
		beego.BConfig.WebConfig.StaticDir["/static"] = "ui/dist/static"
	}
	defer db.Db.Close()
	beego.Run()
}

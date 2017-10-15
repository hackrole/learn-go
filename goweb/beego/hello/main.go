package main

import (
	_ "learn-go/goweb/hello/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


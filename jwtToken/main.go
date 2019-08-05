package main

import (
	_ "jwtToken/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}


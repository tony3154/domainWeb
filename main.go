package main

import (
	_ "DomainDiscoverWeb/models"
	_ "DomainDiscoverWeb/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

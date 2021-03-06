package main

import (
	"ipam-htsc/rest-service/controller"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego"
	// "os"
)

func main() {
	initBeego()
	log.Debugf("IPAM_HTSC is Running")
	beego.Run(":8080")
}

func initBeego() { 
	beego.RESTRouter("/subnet",&controller.SubnetworkController{})
	beego.RESTRouter("/ip",&controller.IPController{})
}

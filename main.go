package main

import (
	setup "dev-event/initialize"
	"dev-event/models"
	"dev-event/routers"

	"github.com/gin-gonic/gin"
)

// @APIVersion 0.1
// @Title Dev Evnet Server
// @Description Dev Event API
func main() {
	r := gin.Default()

	setup.ViperInit()
	models.DBInit()
	defer models.CloseDB()
	setup.CorsInit(r)
	routers.Init(r)

	r.Run(":8090")
}

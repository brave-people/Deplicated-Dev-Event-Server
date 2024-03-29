package main

import (
	setup "dev-event/config"
	"dev-event/docs"
	"dev-event/models"
	"dev-event/routers"

	"github.com/gin-gonic/gin"
)

// @APIVersion 0.1
// @Title Dev Event Server
// @Description Dev Event API
func main() {
	setupServer().Run(":8090")
}

func setupServer() *gin.Engine {
	r := gin.Default()

	setup.ViperInit()
	models.DBInit()
	defer models.CloseDB()
	setup.CorsInit(r)
	routers.Init(r)

	docs.SwaggerInfo.Host = "127.0.0.1:8090"
	r.Run(":8090")
	return r
}

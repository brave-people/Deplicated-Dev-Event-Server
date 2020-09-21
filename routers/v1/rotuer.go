package router

import (
	"dev-event/controllers"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	test := r.Group("/test")
	{
		test.GET("/", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}
	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.Login)
	}
	user := r.Group("/user")
	{
		user.GET("/:userID", controllers.GetMyProfile)
		user.PUT("/:userID", controllers.ModifyProfile)
	}
}

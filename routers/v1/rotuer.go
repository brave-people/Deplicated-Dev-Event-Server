package router

import (
	"dev-event/controllers"
	"dev-event/middlewares"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	test := r.Group("/light")
	{
		test.GET("/", func(c *gin.Context) { c.String(200, "salt") })
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.Login)
	}

	user := r.Group("/users")
	{
		user.GET("/", controllers.GetMyProfile)
		user.PUT("/:email", controllers.ModifyProfile)
	}

	admin := r.Group("/admin")
	{
		users := admin.Group("/users")
		{
			users.GET("/", middlewares.AdminCheck, controllers.AdminGetAllUsers)
			users.GET("/:userID", middlewares.AdminCheck, controllers.AdminGetOneUser)
			users.PUT("/:userID", middlewares.AdminCheck, controllers.AdminUpdateUser)
			users.DELETE("/:userID", middlewares.AdminCheck, controllers.AdminDeleteUser)
		}
		// 	events := r.Group("/events")
		// 	{

		// 	}
	}
}

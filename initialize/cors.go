package init

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsInit(r *gin.Engine) {
	conf := &cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "x-user-auth-token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowOrigins = []string{"*"}
	r.Use(cors.New(*conf))
}

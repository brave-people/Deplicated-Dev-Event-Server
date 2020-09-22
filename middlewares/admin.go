package middlewares

import (
	"dev-event/controllers"
	"dev-event/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminCheck(c *gin.Context) {
	tokenUserID, err := controllers.ExtractTokenID(c.Request)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("No Admin auth"))
		return
	}

	err = models.AdminUserChecker(tokenUserID)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
	}

	return
}

package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var reqUser requests.Login

	err := c.BindJSON(&reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	user, err := models.GetMyProfile(reqUser.Email)
	err = models.VerifyPassword(user.Password, reqUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to auth"})
		return
	}

	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to make JWT"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Token": token})
	return
}

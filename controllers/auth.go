package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login ...
// @Title 로그인
// @Description 로그인
// @Tags Auth
// @Param	body  body     requests.Login  true  "로그인"
// @Success 200   {object} requests.Login
// @Failure 401
// @Failure 500
// @Router /v1/auth/login [Post]
func Login(c *gin.Context) {
	var reqUser requests.Login

	err := c.BindJSON(&reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	user, err := models.GetMyProfileByEmail(reqUser.Email)
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

package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user requests.RegisterUser

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unvaild request"})
		return
	}

	err = models.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create user"})
	return
}

func GetMyProfile(c *gin.Context) {
	userID := c.Param("userID")

	res, err := models.GetMyProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func ModifyProfile(c *gin.Context) {
	userID := c.Param("userID")

	var user requests.RegisterUser
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unvaild request"})
		return
	}

	err = models.ModifyProfile(user, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Update user profile"})
	return
}

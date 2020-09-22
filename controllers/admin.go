package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// AdminGetAllUsers ...
// @Title Get All Users
// @Description Get all users info
// @Tags Admin - Users
// @Success 200 {object} []responses.Users
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/admin/users [Get]
// @Security userAPIKey
func AdminGetAllUsers(c *gin.Context) {
	res, err := models.GetAllUsers()
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invaild ID - record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get minimpas"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AdminGetUser ...
// @Title 회원 정보 조회
// @Description 회원 정보 조회
// @Tags Admin - Users
// @Param	userID	 path	string   true	"userID"
// @Success 200 {object} responses.Users
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/admin/users/:userID [Get]
// @Security userAPIKey
func AdminGetOneUser(c *gin.Context) {
	userID := c.Param("userID")

	res, err := models.GetUser(userID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot get user info"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// AdminUpdateUser ...
// @Title 회원 정보 수정
// @Description 회원 정보 수정
// @Tags Admin - Users
// @Param	userID	 path	string   true	"userID"
// @Success 200 {object} responses.Users
// @Failure 400
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/admin/users/:userID [Put]
// @Security userAPIKey
func AdminUpdateUser(c *gin.Context) {
	userID := c.Param("userID")
	var user requests.Users

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	err = models.AdminUpdateUser(user, userID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved stage infomation."})
}

// AdminDeleteUser ...
// @Title 회원 정보 삭제
// @Description 회원 정보 삭제
// @Tags Admin - Users
// @Param	userID	path	 string            true	   "userID"
// @Success 200     {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/admin/users/:userID [Delete]
// @Security userAPIKey
func AdminDeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	err := models.AdminDeleteUser(userID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	return
}

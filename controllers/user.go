package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 토큰의 ID를 통해 email, param의 email 비교
func VerifyRequestUser(c *gin.Context, userEmail string) (err error) {
	tokenUserID, err := ExtractTokenID(c.Request)
	if err != nil {
		return
	}
	fmt.Println(tokenUserID)

	err = models.VerifyTokenIDAndEmail(tokenUserID, userEmail)
	return
}

// RegisterUser ...
// @Title 회원가입
// @Description 회원가입
// @Tags Auth
// @Param	body  body  requests.RegisterUser true "register user"
// @Success 201 {object} responses.Message
// @Failure 400
// @Failure 500
// @Router /v1/auth/register [Post]
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
}

// GetMyProfile ...
// @Title 나의 프로필 확인
// @Description 나의 프로필 확인
// @Tags Users
// @Param	email	path	string   true	"email"
// @Success 200 {object} responses.Users
// @Failure 401
// @Failure 500
// @Router /v1/users [Get]
// @Security userAPIKey
func GetMyProfile(c *gin.Context) {
	tokenUserID, err := ExtractTokenID(c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token error"})
		return
	}

	res, err := models.GetMyProfileByUserID(tokenUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// ModifyProfile ...
// @Title 프로필 수정
// @Description 프로필 수정
// @Tags Users
// @Param	email	path	 string             true	"email"
// @Success 200   {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/users/:email [Put]
// @Security userAPIKey
func ModifyProfile(c *gin.Context) {
	email := c.Param("email")

	err := VerifyRequestUser(c, email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	var user requests.RegisterUser
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unvaild request"})
		return
	}

	err = models.ModifyProfile(user, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"message": "Update user profile"})
}

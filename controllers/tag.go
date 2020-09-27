package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetAllTags ...
// @Title 모든 태그 조회
// @Description 모든 태그 조회
// @Tags Tags
// @Success 200 {object} []responses.Tag
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/tags [Get]
// @Security userAPIKey
func GetAllTags(c *gin.Context) {
	res, err := models.GetAllTags()
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rcord not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get tag"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// DeleteTag ...
// @Title 태그 삭제
// @Description 태그 삭제
// @Tags Tags
// @Success 200 {object} []responses.Message
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/tags [Get]
// @Security userAPIKey
func DeleteTag(c *gin.Context) {
	tagID := c.Param("tagID")

	err := models.DeleteTag(tagID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to delete tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}

// UpdateTag ...
// @Title 태그 수정
// @Description 태그 수정
// @Tags Tours
// @Param   tagID	path	string  true	"tagID"
// @Success 200 {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /admin/tags/:tagID [Put]
func UpdateTag(c *gin.Context) {
	tagID := c.Param("tagID")
	var tag requests.Tag

	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	err = models.UpdateTag(tag, tagID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Invaild ID"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save tag"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved tag"})
}

// CreateTag ...
// @Title 태그 생성
// @Description 태그 생성
// @Tags Tours
// @Param   tagID	path	string  true	"tagID"
// @Success 200 {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /admin/tags [Post]
func CreateTag(c *gin.Context) {
	var tag requests.Tag
	err := c.BindJSON(&tag)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	res, err := models.CreateTag(tag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to create stage"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": res})
}

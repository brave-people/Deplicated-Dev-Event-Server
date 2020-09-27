package controllers

import (
	"dev-event/models"
	"dev-event/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GetOneEvent ...
// @Title 이벤트 정보 조회
// @Description 이벤트 정보 조회
// @Tags Event
// @Param	eventID	 path	string   true	"eventID"
// @Success 200 {object} models.Event
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/events/:eventID [Get]
// @Security userAPIKey
func GetOneEvent(c *gin.Context) {
	eventID := c.Param("eventID")

	res, err := models.GetOneEvent(eventID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rcord not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get event"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// UpdateEvent ...
// @Title 이벤트 수정
// @Description 이벤트 수정
// @Tags Event
// @Param	body     body   requests.Event true  "Save event"
// @Param	eventID	 path	string         true	 "eventID"
// @Success 200 {object} models.Event
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/events/:eventID [Put]
// @Security userAPIKey
func UpdateEvent(c *gin.Context) {
	eventID := c.Param("eventID")
	var event requests.Event

	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}
	err = models.SaveEvent(event, eventID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to save event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully saved event"})
}

// DeleteEvent ...
// @Title 이벤트 삭제
// @Description 이벤트 삭제
// @Tags Event
// @Param	eventID	 path	string   true	"eventID"
// @Success 200 {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /v1/admin/events/:eventID [Get]
// @Security userAPIKey
func DeleteEvent(c *gin.Context) {
	eventID := c.Param("eventID")

	err := models.DeleteEvent(eventID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "No record found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to delete evenet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

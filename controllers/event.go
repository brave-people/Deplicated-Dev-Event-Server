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
// @Tags Admin - Events
// @Param	eventID	 path	string   true	"eventID"
// @Success 200 {object} models.Event
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /admin/events/:eventID [Get]
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
// @Tags Admin - Events
// @Param	body     body   requests.Event true  "Save event"
// @Param	eventID	 path	string         true	 "eventID"
// @Success 200 {object} models.Event
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /admin/events/:eventID [Put]
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
// @Tags Admin - Events
// @Param	eventID	 path	string   true	"eventID"
// @Success 200 {object} responses.Message
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /admin/events/:eventID [Delete]
// @Security userAPIKey
func DeleteEvent(c *gin.Context) {
	eventID := c.Param("eventID")

	err := models.DeleteEvent(eventID)
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Record not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to delete evenet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

// GetEvent ...
// @Title 이벤트 가져오기 (년, 월)
// @Description 이벤트 가져오기 (년, 월)
// @Tags Events
// @Param	year	 path	string   true	"year"
// @Param	month	 path	string   true	"month"
// @Success 200 {object} []responses.Event
// @Failure 401
// @Failure 404
// @Failure 500
// @Router /events/:year/:month [Get]
func GetYYMMEvents(c *gin.Context) {
	year := c.Param("year")
	month := c.Param("month")

	res, err := models.GetYYMMEvents(year, month)
	if len(res) == 0 || err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"message": "Rcord not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to get event"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateEvent ...
// @Title 새 이벤트 생성
// @Description 새 이벤트 생성
// @Tags Admin - Events
// @Param	body     body     requests.Event true  "create event"
// @Success 201      {object} responses.Message
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /admin/events [Post]
// @Security userAPIKey
func CreateEvent(c *gin.Context) {
	var event requests.Event

	err := c.BindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Fail to bind request"})
		return
	}

	res, err := models.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Fail to create event"})
		return
	}

	c.JSON(http.StatusCreated, res)
}

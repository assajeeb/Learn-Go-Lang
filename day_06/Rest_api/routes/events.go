package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, event)

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"mess": err})
	}

	context.JSON(http.StatusOK, events)
}

func createEvents(context *gin.Context) {
	var event models.Event
	erro := context.ShouldBindJSON(&event)
	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Failed to Parse request"})
		return
	}
	userid := context.GetInt64("userid")
	event.UserId = userid
	erro = event.Save()
	if erro != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Failed to save request"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created Successfully", "event": event})

}

func updateEvents(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	userid := context.GetInt64("userid")
	if userid != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to update this", "eventuserid": event.UserId, "userid": userid})
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to bind json"})
		return
	}

	updatedEvent.Id = event.Id

	err = updatedEvent.Update()
	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updateded successfully"})

}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad parameter"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "can not fetch the event"})
		return
	}
	userid := context.GetInt64("userid")
	if userid != event.UserId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "you are not authorized to update this"})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to delete event"})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "Event deleted successfuly!"})

}

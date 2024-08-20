package routes

import (
	"example/com/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)

}

func getEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var events models.Event
	err := context.ShouldBindJSON(&events)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	events.ID = 1
	events.UserID = 1

	err = events.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "event": events})
}

func updateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid event id"})
		return
	}

	_, err = models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event by id"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}

	updatedEvent.ID = eventID
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated", "event": updatedEvent})

}
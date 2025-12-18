package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/event/:id", getEvent)
	server.PUT("/event/:id", updateEvents)
	server.DELETE("/event/:id", deleteEvent)
	server.POST("/signup", SignUpUser)
	server.POST("/login", login)
}

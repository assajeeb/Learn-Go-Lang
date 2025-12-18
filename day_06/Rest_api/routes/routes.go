package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authorization)
	authenticated.POST("/events", createEvents)
	authenticated.PUT("/event/:id", updateEvents)
	authenticated.DELETE("/event/:id", deleteEvent)

	server.POST("/signup", SignUpUser)
	server.POST("/login", login)
}

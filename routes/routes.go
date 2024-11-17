package routes

import (
	"bitcoin-rate/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes is used for endpoints registration
func RegisterRoutes(server *gin.Engine) {
	// Get request for receiving current BTC to UAH rate
	// Returns json with current rate of BTS
	server.GET("/rate", controllers.GetRate)

	// Subscribe given email on mailing list
	server.POST("/subscribe", controllers.Subscribe)
	// Send emails with current rate on all subscribed addresses
	server.POST("/sendEmails", controllers.SendEmails)
	// Get all emails that subscribed to mailing list
	server.POST("/getAllEmails", controllers.GetAllEmails)
}

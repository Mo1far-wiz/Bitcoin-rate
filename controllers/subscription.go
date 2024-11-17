package controllers

import (
	"bitcoin-rate/models"
	emailer "bitcoin-rate/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Subscribe(context *gin.Context) {
	var email models.Email

	err := context.ShouldBindJSON(&email)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"message": "could not parse request data."})
		return
	}

	err = email.Save()
	if err != nil {
		log.Println("Error: ", err.Error())
		context.JSON(http.StatusConflict, gin.H{"message": "could not save email."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "email saved successfully.", "email": email})
}

func GetAllEmails(context *gin.Context) {
	emails, err := models.GetAllEmails()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch emails."})
		return
	}

	context.JSON(http.StatusOK, emails)
}

func SendEmails(context *gin.Context) {
	emails, err := models.GetAllEmails()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch emails."})
		return
	}

	rate, err := rs.GetBTCRate()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	for _, email := range emails {
		go emailer.SendEmail(
			email.Email,
			"Current BTC rate",
			fmt.Sprintf("Hello!\nCurrent BTC rate is %f UAH.", rate),
		)
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
}

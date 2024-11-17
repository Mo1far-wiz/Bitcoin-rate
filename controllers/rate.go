package controllers

import (
	"bitcoin-rate/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rs *adapters.RemoteService = &adapters.RemoteService{
	Remote: &adapters.CoinbaseApi{},
}

// Using RemoteService adapter gets current rate of btc
func GetRate(context *gin.Context) {
	rate, err := rs.GetBTCRate()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"rate": rate})
}

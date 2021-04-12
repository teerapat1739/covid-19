package transport

import (
	"covid-19/business/covid"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CovidHTTPHandler struct {
	covidService covid.Service
}

func NewCovidHTTPHandler(router *gin.Engine, covidService covid.Service) {
	// handler := CovidHTTPHandler{covidService}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

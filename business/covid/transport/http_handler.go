package transport

import (
	"covid-19/business/covid"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CovidHTTPHandler struct {
	covidService covid.Service
}

// ok  	covid-19/business/covid/transport	0.389s	coverage: 100.0% of statements

func NewCovidHTTPHandler(covidService covid.Service) *gin.Engine {
	router := gin.Default()

	h := CovidHTTPHandler{covidService}

	router.GET("/covid/summary", h.GetCovidSummary)
	return router
}

func (h *CovidHTTPHandler) GetCovidSummary(c *gin.Context) {
	res, err := h.covidService.CovidSummary()
	log.Println("errrrr", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

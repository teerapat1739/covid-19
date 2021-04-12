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
	h := CovidHTTPHandler{covidService}

	router.GET("/covid/summary", h.GetCovidSummary)
}

func (h *CovidHTTPHandler) GetCovidSummary(c *gin.Context) {
	res, err := h.covidService.CovidSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, res)
}

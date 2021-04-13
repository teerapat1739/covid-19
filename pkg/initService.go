package pkg

import (
	"net/http"
	"time"

	covidrepo "covid-19/business/covid/repository"
	covidsrc "covid-19/business/covid/service"
	covidhdlr "covid-19/business/covid/transport"

	"github.com/gin-gonic/gin"
)

func InintService() *gin.Engine {
	router := gin.Default()
	client := &http.Client{
		Timeout: time.Duration(30 * time.Second),
	}

	// covidRepo := covidrepo.NewCovidRepository(client, os.Getenv("BASE_COVID_URL"))
	covidRepo := covidrepo.NewCovidRepository(client, "https://covid19.th-stat.com/api")
	covidSrc := covidsrc.NewCovidService(covidRepo)
	covidhdlr.NewCovidHTTPHandler(router, covidSrc)

	return router
}

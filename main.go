package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"

	covidrepo "covid-19/business/covid/repository"
	covidsrc "covid-19/business/covid/service"
	covidhdlr "covid-19/business/covid/transport"
)

func init() {
	gotenv.Load()
}

func main() {

	router := gin.Default()

	covidRepo := covidrepo.NewCovidRepository(os.Getenv("BASE_COVID_URL"))
	covidSrc := covidsrc.NewCovidService(covidRepo)
	covidhdlr.NewCovidHTTPHandler(router, covidSrc)

	router.Run()

}

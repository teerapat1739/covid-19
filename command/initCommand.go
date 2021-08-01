package command

import (
	covidrepo "covid-19/business/covid/repository"
	"covid-19/pkg"

	covidsrc "covid-19/business/covid/service"
	covidhdlr "covid-19/business/covid/transport"
	"os"
)

func InintCommand() {

	client := pkg.NewHttpClient()
	covidRepo := covidrepo.NewCovidRepository(client, os.Getenv("BASE_COVID_URL"))
	covidSrc := covidsrc.NewCovidService(covidRepo)
	router := covidhdlr.NewCovidHTTPHandler(covidSrc)
	router.Run(":8080")

}

package covid

import "covid-19/business/covid/model"

type Repository interface {
	GetCovidCases() (model.ReponseCovidApiCases, error)
}

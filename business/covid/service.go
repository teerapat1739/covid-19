package covid

import "covid-19/business/covid/model"

type Service interface {
	CovidSummary() (model.Summary, error)
}

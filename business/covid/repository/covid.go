package repository

import (
	"covid-19/business/covid"
)

type covidRepository struct {
	baseCovidUrl string
}

func NewCovidRepository(baseCovidUrl string) covid.Repository {
	return &covidRepository{
		baseCovidUrl: baseCovidUrl}
}

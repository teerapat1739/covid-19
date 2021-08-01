package service

import "covid-19/business/covid"

type covidService struct {
	covidRepo covid.Repository
}

func NewCovidService(covidRepo covid.Repository) covid.Service {
	return &covidService{covidRepo: covidRepo}
}

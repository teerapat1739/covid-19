package service

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/stretchr/testify/mock"
)

type covidService struct {
	covidRepo covid.Repository
}

func NewCovidService(covidRepo covid.Repository) covid.Service {
	return &covidService{covidRepo: covidRepo}
}

func (c *covidService) CovidSummary() (model.Summary, error) {
	res, err := c.covidRepo.GetCovidOpenCases()
	if err != nil {
		return model.Summary{}, err
	}
	provinceSum := make(map[string]float64)
	ageSum := make(map[string]float64)

	for _, val := range res.Data {
		if _, ok := provinceSum[val.Province]; ok {
			provinceSum[val.Province] += 1
		} else {
			provinceSum[val.Province] = 0
		}

		ageFiledName := ""
		if val.Age == nil {
			ageFiledName = "N/A"
		} else {
			age := val.Age.(float64)
			if age <= 30 {
				ageFiledName = "0-30"
			} else if age >= 31 && age <= 60 {
				ageFiledName = "31-60"
			} else {
				ageFiledName = "61+"
			}
		}

		if _, ok := ageSum[ageFiledName]; ok {
			ageSum[ageFiledName] += 1
		} else {
			ageSum[ageFiledName] = 0
		}
	}

	var m model.Summary
	m.Province = provinceSum
	m.AgeGroup = ageSum
	return m, err
}

type CovidRepomock struct {
	mock.Mock
}

func (m *CovidRepomock) GetCovidOpenCases() (model.ReponseApiOpenCases, error) {
	jsonFile, _ := os.Open("covid.json")
	var c model.ReponseApiOpenCases

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &c)
	return c, nil
}

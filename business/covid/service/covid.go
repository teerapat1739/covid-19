package service

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/mock"
)

type covidService struct {
	covidRepo covid.Repository
}

func NewCovidService(covidRepo covid.Repository) covid.Service {
	return &covidService{covidRepo: covidRepo}
}

func (c *covidService) CovidSummary() ([]model.ResponseApi, error) {
	res, err := c.covidRepo.GetCovidOpenCases()
	if err != nil {
		return nil, err
	}
	m := make(map[int]model.ResponseApi)

	for _, val := range res.Data {
		if v, ok := m[val.Provinceid]; ok {
			resApi := v
			resApi.Totaluser += 1
			if val.Age == nil {
				resApi.Groupage.Agedontknow += 1
			} else {
				age := val.Age.(float64)
				if age <= 30 {
					resApi.Groupage.Agelessthan30 += 1
				} else if age > 31 && age <= 60 {
					resApi.Groupage.Agebetween31And60 += 1
				} else {
					resApi.Groupage.Agemorethan60 += 1
				}
			}
			m[val.Provinceid] = resApi

		} else {
			var resApi model.ResponseApi
			resApi.Proince = val.Province
			resApi.Provinceid = val.Provinceid
			resApi.Totaluser = 1
			if val.Age == nil {
				resApi.Groupage.Agedontknow += 1
			} else {
				age := val.Age.(float64)
				if age <= 30 {
					resApi.Groupage.Agelessthan30 += 1
				} else if age > 31 && age <= 60 {
					resApi.Groupage.Agebetween31And60 += 1
				} else {
					resApi.Groupage.Agemorethan60 += 1
				}
			}

			m[val.Provinceid] = resApi
		}

	}

	var data []model.ResponseApi
	for _, element := range m {
		var d model.ResponseApi

		mapstructure.Decode(element, &d)
		data = append(data, d)
	}
	return data, err
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

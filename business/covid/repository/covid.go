package repository

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"covid-19/pkg"
	"reflect"

	"encoding/json"
)

type covidRepository struct {
	client       pkg.HTTPClient
	baseCovidUrl string
}

// ok  	covid-19/business/covid/repository	0.811s	coverage: 91.7% of statements

func NewCovidRepository(client pkg.HTTPClient, baseCovidUrl string) covid.Repository {
	return &covidRepository{
		client:       client,
		baseCovidUrl: baseCovidUrl}

}

func (c *covidRepository) GetCovidCases() (model.ReponseCovidApiCases, error) {
	var m model.ReponseCovidApiCases
	url := c.baseCovidUrl + "devinterview/covid-cases.json"
	resp, err := c.client.Get(url)
	if err != nil {
		return m, err
	}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if reflect.DeepEqual(m, model.ReponseCovidApiCases{}) || err != nil {
		return m, err
	}
	defer resp.Body.Close()

	return m, nil

}

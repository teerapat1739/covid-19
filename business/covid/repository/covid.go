package repository

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"encoding/json"
	"net/http"
)

type covidRepository struct {
	client       *http.Client
	baseCovidUrl string
}

func NewCovidRepository(client *http.Client, baseCovidUrl string) covid.Repository {
	return &covidRepository{
		client:       client,
		baseCovidUrl: baseCovidUrl}

}

func (c *covidRepository) GetCovidOpenCases() (model.ReponseApiOpenCases, error) {
	var m model.ReponseApiOpenCases
	resp, err := c.client.Get(c.baseCovidUrl + "/open/cases")
	if err != nil {
		return m, err
	}
	json.NewDecoder(resp.Body).Decode(&m)
	defer resp.Body.Close()

	return m, nil
}

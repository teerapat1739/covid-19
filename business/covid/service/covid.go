package service

import (
	"covid-19/business/covid/model"
)

// ok  	covid-19/business/covid/service	0.340s	coverage: 97.1% of statements
func (c *covidService) CovidSummary() (model.Summary, error) {
	res, err := c.covidRepo.GetCovidCases()
	if err != nil {
		return model.Summary{}, err
	}
	provinceSum := make(map[string]float64)
	ageSum := make(map[string]float64)

	for _, val := range res.Data {
		provinceSum = c.calProvince(val, provinceSum)

		ageSum = c.calAge(val, ageSum)
	}

	var m model.Summary
	m.Province = provinceSum
	m.AgeGroup = ageSum
	return m, err
}

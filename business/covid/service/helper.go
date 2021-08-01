package service

import (
	"covid-19/business/covid/model"
)

func (c *covidService) calProvince(data model.CovidDetail, provinceSum map[string]float64) map[string]float64 {
	if data.Province == "" {
		provinceSum["N/A"] += 1
	} else if _, ok := provinceSum[data.Province]; ok {
		provinceSum[data.Province] += 1
	} else {
		provinceSum[data.Province] = 1
	}
	return provinceSum
}

func (c *covidService) calAge(data model.CovidDetail, ageSum map[string]float64) map[string]float64 {
	ageFiledName := ""

	if data.Age == nil {
		ageFiledName = "N/A"
	} else if *data.Age < 0 {
		ageFiledName = "N/A"
	} else {
		age := *data.Age
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
		ageSum[ageFiledName] = 1
	}
	return ageSum
}

package service_test

import (
	"covid-19/business/covid/model"
	"covid-19/business/covid/service"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCovidSummary(t *testing.T) {
	jsonFile, _ := os.Open("covid.json")
	var c model.ReponseApiOpenCases

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &c)
	covidSrcMock := new(service.CovidRepomock)
	covidSrcMock.GetCovidOpenCases()
	covidSrcMock.On("GetCovidOpenCases", nil).Return(c, nil)

	covidSrc := service.NewCovidService(covidSrcMock)

	actual, err := covidSrc.CovidSummary()

	expected := []model.ResponseApi{}

	expected = append(expected, model.ResponseApi{
		Proince:    "สมุทรสาคร",
		Provinceid: 62,
		Totaluser:  43,
		Groupage: model.Groupage{
			Agelessthan30:     1,
			Agebetween31And60: 12,
			Agemorethan60:     5,
			Agedontknow:       25,
		},
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

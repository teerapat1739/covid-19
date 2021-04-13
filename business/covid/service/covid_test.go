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

	var expected model.Summary
	expected.Province = map[string]float64{"สมุทรสาคร": float64(42)}
	expected.AgeGroup = map[string]float64{"0-30": float64(0), "31-60": float64(11), "61+": float64(4), "N/A": float64(24)}

	assert.Equal(t, nil, err)
	assert.Equal(t, expected.Province, actual.Province)
	assert.Equal(t, expected.AgeGroup, actual.AgeGroup)
}

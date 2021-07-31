package service

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"reflect"
	"testing"
)

func Test_covidService_CovidSummary(t *testing.T) {
	type fields struct {
		covidRepo covid.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.Summary
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidService{
				covidRepo: tt.fields.covidRepo,
			}
			got, err := c.CovidSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("covidService.CovidSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidService.CovidSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

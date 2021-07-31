package repository

import (
	"covid-19/business/covid/model"
	"net/http"
	"reflect"
	"testing"
)

func Test_covidRepository_GetCovidCases(t *testing.T) {
	type fields struct {
		client       *http.Client
		baseCovidUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		want    model.ReponseApiOpenCases
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidRepository{
				client:       tt.fields.client,
				baseCovidUrl: tt.fields.baseCovidUrl,
			}
			got, err := c.GetCovidCases()
			if (err != nil) != tt.wantErr {
				t.Errorf("covidRepository.GetCovidCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidRepository.GetCovidCases() = %v, want %v", got, tt.want)
			}
		})
	}
}

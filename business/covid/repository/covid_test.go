package repository

import (
	"bytes"
	"covid-19/business/covid/model"
	"covid-19/pkg"
	mock_covid "covid-19/test"
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

const jsonTest = `
{
    "Data": [
        {
            "Age": 33,
            "Province": "Trang"
        },
        {
            "Age": 51,
            "Province": "Suphan Buri"
        }
    ]
}
`

func Test_covidRepository_GetCovidCases(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mClient := mock_covid.NewMockHTTPClient(ctrl)

	type fields struct {
		client       pkg.HTTPClient
		baseCovidUrl string
	}
	tests := []struct {
		name          string
		fields        fields
		mockResponses *http.Response
		want          model.ReponseCovidApiCases
		wantErr       bool
	}{
		{
			name: "test_getCovidCases_api_success",
			fields: fields{
				client:       mClient,
				baseCovidUrl: "",
			},
			mockResponses: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(jsonTest))),
			},
			want: model.ReponseCovidApiCases{
				Data: []model.CovidDetail{
					{
						Age:      pkg.GetIntPointer(33),
						Province: "Trang",
					},
					{
						Age:      pkg.GetIntPointer(51),
						Province: "Suphan Buri",
					},
				},
			},
		},
		{
			name: "test_getCovidCases_api_fail",
			fields: fields{
				client:       mClient,
				baseCovidUrl: "",
			},
			mockResponses: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(``))),
			},
			want:    model.ReponseCovidApiCases{},
			wantErr: true,
		},
		{
			name: "test_getCovidCases_api_success_but_empty_response_body",
			fields: fields{
				client:       mClient,
				baseCovidUrl: "",
			},
			mockResponses: &http.Response{
				StatusCode: 500,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"vaccine": "sinnovac"}`))),
			},
			want:    model.ReponseCovidApiCases{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidRepository{
				client:       tt.fields.client,
				baseCovidUrl: tt.fields.baseCovidUrl,
			}
			var err error
			if tt.mockResponses.StatusCode != 200 {
				err = errors.New("something wrong")
			}
			mClient.EXPECT().Get(gomock.Any()).Return(tt.mockResponses, err)

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

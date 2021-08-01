package service

import (
	"covid-19/business/covid"
	"covid-19/business/covid/model"
	"covid-19/pkg"
	"errors"
	"log"
	"reflect"
	"testing"

	mock_covid "covid-19/test"

	"github.com/golang/mock/gomock"
)

func Test_covidService_calProvince(t *testing.T) {
	type fields struct {
		covidRepo covid.Repository
	}
	type args struct {
		data        model.CovidDetail
		provinceSum map[string]float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]float64
	}{

		{
			name: "test_calProvince_1: When feild provice is empty",
			args: args{
				data: model.CovidDetail{
					Province: "",
				},
				provinceSum: map[string]float64{},
			},
			want: map[string]float64{
				"N/A": 1,
			},
		},
		{
			name: "test_calProvince_2: When provice does not exits in map and map is empty",
			args: args{
				data: model.CovidDetail{
					Province: "wongnai",
				},
				provinceSum: map[string]float64{},
			},
			want: map[string]float64{
				"wongnai": 1,
			},
		},
		{
			name: "test_calProvince_3: When provice does not exits in map and map is not empty",
			args: args{
				data: model.CovidDetail{
					Province: "wongnai",
				},
				provinceSum: map[string]float64{
					"lineman": 3,
				},
			},
			want: map[string]float64{
				"wongnai": 1,
				"lineman": 3,
			},
		},
		{
			name: "test_calProvince_4: When provice does  exits in map",
			args: args{
				data: model.CovidDetail{
					Province: "wongnai",
				},
				provinceSum: map[string]float64{
					"lineman": 3,
					"wongnai": 2,
				},
			},
			want: map[string]float64{
				"wongnai": 3,
				"lineman": 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidService{
				covidRepo: tt.fields.covidRepo,
			}
			if got := c.calProvince(tt.args.data, tt.args.provinceSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidService.calProvince() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_covidService_calAge(t *testing.T) {
	type fields struct {
		covidRepo covid.Repository
	}
	type args struct {
		data   model.CovidDetail
		ageSum map[string]float64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]float64
	}{
		{
			name: "test_calAge_1: When age is empty",
			args: args{
				data:   model.CovidDetail{},
				ageSum: map[string]float64{},
			},
			want: map[string]float64{
				"N/A": 1,
			},
		},
		{
			name: "test_calAge_2: When age less than 0",
			args: args{
				data: model.CovidDetail{
					Age: pkg.GetIntPointer(-2),
				},
				ageSum: map[string]float64{},
			},
			want: map[string]float64{
				"N/A": 1,
			},
		},
		{
			name: "test_calAge_3: When age <= 30",
			args: args{
				data: model.CovidDetail{
					Age: pkg.GetIntPointer(29),
				},
				ageSum: map[string]float64{},
			},
			want: map[string]float64{
				"0-30": 1,
			},
		},
		{
			name: "test_calAge_4: When age between 31 and 60",
			args: args{
				data: model.CovidDetail{
					Age: pkg.GetIntPointer(39),
				},
				ageSum: map[string]float64{},
			},
			want: map[string]float64{
				"31-60": 1,
			},
		},
		{
			name: "test_calAge_4: When age more than 60",
			args: args{
				data: model.CovidDetail{
					Age: pkg.GetIntPointer(69),
				},
				ageSum: map[string]float64{},
			},
			want: map[string]float64{
				"61+": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidService{
				covidRepo: tt.fields.covidRepo,
			}
			if got := c.calAge(tt.args.data, tt.args.ageSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidService.calAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_covidService_CovidSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mRepo := mock_covid.NewMockRepository(ctrl)
	type fields struct {
		covidRepo covid.Repository
	}

	type returnData struct {
		m   model.ReponseCovidApiCases
		err error
	}
	tests := []struct {
		name       string
		fields     fields
		returnData returnData
		want       model.Summary
		wantErr    bool
	}{
		{
			name: "test_CovidSummary_1: When no data",
			fields: fields{
				covidRepo: mRepo,
			},
			returnData: returnData{
				m:   model.ReponseCovidApiCases{},
				err: nil,
			},
			want: model.Summary{
				Province: map[string]float64{},
				AgeGroup: map[string]float64{},
			},
		},
		{
			name: "test_CovidSummary_2: When repo have error",
			fields: fields{
				covidRepo: mRepo,
			},
			returnData: returnData{
				m:   model.ReponseCovidApiCases{},
				err: errors.New("error"),
			},
			want:    model.Summary{},
			wantErr: true,
		},
		{
			name: "test_CovidSummary_2: When repo have return data",
			fields: fields{
				covidRepo: mRepo,
			},
			returnData: returnData{
				m: model.ReponseCovidApiCases{
					Data: []model.CovidDetail{
						{
							Province: "wong",
							Age:      pkg.GetIntPointer(1),
						},
						{
							Province: "nai",
							Age:      pkg.GetIntPointer(33),
						},
						{
							Province: "line",
							Age:      pkg.GetIntPointer(64),
						},
						{
							Province: "man",
							Age:      pkg.GetIntPointer(-1),
						},
						{
							Province: "prayut",
						},
						{
							Province: "",
							Age:      pkg.GetIntPointer(33),
						},
						{
							Province: "",
							Age:      pkg.GetIntPointer(0),
						},
						{
							Province: "nai",
							Age:      pkg.GetIntPointer(21),
						},
						{
							Province: "man",
						},
					},
				},
				err: nil,
			},
			want: model.Summary{
				Province: map[string]float64{
					"wong":   1,
					"nai":    2,
					"line":   1,
					"man":    2,
					"prayut": 1,
					"N/A":    2,
				},
				AgeGroup: map[string]float64{
					"N/A":   3,
					"0-30":  3,
					"31-60": 2,
					"61+":   1,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &covidService{
				covidRepo: tt.fields.covidRepo,
			}
			// var m model.ReponseCovidApiCases
			mRepo.EXPECT().GetCovidCases().Return(tt.returnData.m, tt.returnData.err)
			got, err := c.CovidSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("covidService.CovidSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Println(got)
			log.Println(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidService.CovidSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

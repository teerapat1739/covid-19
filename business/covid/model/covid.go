package model

// ReponseApiSummary represent reponse api form wongnai

type ReponseCovidApiCases struct {
	Data []CovidDetail `json:"Data"`
}

type CovidDetail struct {
	Confirmdate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            *int        `json:"Age"`
	Gender         string      `json:"Gender"`
	Genderen       string      `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	Nationen       string      `json:"NationEn"`
	Province       string      `json:"Province"`
	Provinceid     int         `json:"ProvinceId"`
	District       interface{} `json:"District"`
	Provinceen     string      `json:"ProvinceEn"`
	Statquarantine int         `json:"StatQuarantine"`
}

type Summary struct {
	Province map[string]float64 `json:"Province"`
	AgeGroup map[string]float64 `json:"AgeGroup"`
}

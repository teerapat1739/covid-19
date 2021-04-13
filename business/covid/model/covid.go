package model

// ReponseApiSummary represent reponse api form open/cases
type ReponseApiOpenCases struct {
	Data []CovidDetail
}

type CovidDetail struct {
	Confirmdate    string      `json:"ConfirmDate"`
	No             string      `json:"No"`
	Age            interface{} `json:"Age"`
	Gender         string      `json:"Gender"`
	Genderen       string      `json:"GenderEn"`
	Nation         string      `json:"Nation"`
	Nationen       string      `json:"NationEn"`
	Province       string      `json:"Province"`
	Provinceid     int         `json:"ProvinceId"`
	District       string      `json:"District"`
	Provinceen     string      `json:"ProvinceEn"`
	Detail         interface{} `json:"Detail"`
	Statquarantine int         `json:"StatQuarantine"`
}

type Summary struct {
	Province map[string]float64 `json:"Province"`
	AgeGroup map[string]float64 `json:"AgeGroup"`
}

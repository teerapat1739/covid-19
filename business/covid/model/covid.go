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

type ResponseApi struct {
	Proince    string   `json:"Proince"`
	Provinceid int      `json:"ProvinceId"`
	Totaluser  int      `json:"totalUser"`
	Groupage   Groupage `json:"GroupAge"`
}

type Groupage struct {
	Agelessthan30     int `json:"AgeLessThan30"`
	Agebetween31And60 int `json:"AgeBetween31And60"`
	Agemorethan60     int `json:"AgeMoreThan60"`
	Agedontknow       int `json:"AgeDontKnow"`
}

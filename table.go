package main

type MainTable struct {
	Language         string `json:"language"`
	DeclareTime      string `json:"declareTime"`
	PersonNo         string `json:"personNo"`
	PersonName       string `json:"personName"`
	Sex              string `json:"sex"`
	ProfessionName   string `json:"professionName"`
	CollegeName      string `json:"collegeName"`
	PhoneArea        string `json:"phoneArea"`
	Phone            string `json:"phone"`
	AssistantName    string `json:"assistantName"`
	AssistantNo      string `json:"assistantNo"`
	ClassName        string `json:"className"`
	Linkman          string `json:"linkman"`
	LinkmanPhoneArea string `json:"linkmanPhoneArea"`
	LinkmanPhone     string `json:"linkmanPhone"`
	PersonHealth     string `json:"personHealth"`
	Temperature      string `json:"temperature"`
	PersonHealth2    string `json:"personHealth2"`
	SchoolC1         string `json:"schoolC1"`
	CurrentArea      string `json:"currentArea"`
	PersonC4         string `json:"personC4"`
	OtherC4          string `json:"otherC4"`
	IsPass14C1       string `json:"isPass14C1"`
	IsPass14C2       string `json:"isPass14C2"`
	IsPass14C3       string `json:"isPass14C3"`
	PassAreaC2       string `json:"passAreaC2"`
	PassAreaC3       string `json:"passAreaC3"`
	PassAreaC4       string `json:"passAreaC4"`
}

type SecondTable struct {
	Other1 string `json:"other1"`
	Other3 string `json:"other3"`
	Other4 string `json:"other4"`
	Other5 string `json:"other5"`
	Other6 string `json:"other6"`
	Other7 string `json:"other7"`
	Other8 string `json:"other8"`
	Other9 string `json:"other9"`
}

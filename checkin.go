package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CheckinReqData struct {
	JnuID       string      `json:"jnuid"`
	MainTable   MainTable   `json:"mainTable"`
	SecondTable SecondTable `json:"secondTable"`
}

type CheckinRespData struct {
	Meta struct {
		Msg     string `json:"msg"`
		Code    int    `json:"code"`
		Success bool   `json:"success"`
	} `json:"meta"`
}

// Checkin 执行打卡
func Checkin(jnuID string, mainTable MainTable, secondTable SecondTable) (bool, string) {
	reqData := &CheckinReqData{
		JnuID:       jnuID,
		MainTable:   mainTable,
		SecondTable: secondTable,
	}
	reqBody, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", "https://stuhealth.jnu.edu.cn/api/write/main", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var respData CheckinRespData
	json.Unmarshal(respBody, &respData)

	return respData.Meta.Success, respData.Meta.Msg
}

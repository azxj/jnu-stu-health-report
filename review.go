package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type ReviewReqData struct {
	ID    string `json:"id"`
	JnuID string `json:"jnuid"`
}

type ReviewRespData struct {
	Data struct {
		MainTable   MainTable   `json:"mainTable"`
		SecondTable SecondTable `json:"secondTable"`
	} `json:"data"`
}

// Review 获取用户最近一次打卡的所有详细信息，以两个 table 的形式进行返回
func Review(id int, jnuID string) (mainTable MainTable, secondTable SecondTable) {
	reqData := &ReviewReqData{
		ID:    strconv.Itoa(id),
		JnuID: jnuID,
	}
	reqBody, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", "https://stuhealth.jnu.edu.cn/api/user/review", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var respData ReviewRespData
	json.Unmarshal(respBody, &respData)

	mainTable = respData.Data.MainTable
	mainTable.DeclareTime = time.Now().Format("2006-01-02")
	secondTable = respData.Data.SecondTable
	return
}

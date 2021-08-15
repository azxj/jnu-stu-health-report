package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ListReqData struct {
	JnuID string `json:"jnuid"`
}

type ListRespData struct {
	Data struct {
		CheckinInfo []struct {
			ID int `json:"id"`
		} `json:"checkinInfo"`
	} `json:"data"`
}

// List 列出用户的所有打卡记录，返回最近一次打卡的 id
func List(jnuID string) (firstID int) {
	reqData := &ListReqData{
		JnuID: jnuID,
	}
	reqBody, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", "https://stuhealth.jnu.edu.cn/api/user/stucheckin", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var respData ListRespData
	json.Unmarshal(respBody, &respData)

	firstID = 0
	for i := 0; firstID == 0; i++ {
		firstID = respData.Data.CheckinInfo[i].ID
	}
	return
}

package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LoginReqData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRespData struct {
	Data struct {
		JnuID1 string `json:"jnuid"`
		JnuID2 string `json:"jnuId"`
	} `json:"data"`
}

// Login 进行登录，返回用户的唯一标识 jnuid
func Login(username, password string) (jnuID string) {
	reqData := &LoginReqData{
		Username: username,
		Password: password,
	}
	reqBody, _ := json.Marshal(reqData)
	req, _ := http.NewRequest("POST", "https://stuhealth.jnu.edu.cn/api/user/login", bytes.NewReader(reqBody))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)
	respBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var respData LoginRespData
	json.Unmarshal(respBody, &respData)

	jnuID = respData.Data.JnuID1
	return
}

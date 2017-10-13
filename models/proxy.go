package models

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"restapp/common"
	"github.com/beego/bee/logger"
)

type Proxy struct {
	Id       int64
	UserName string
	Address  string
	Phone    string
}

type ProxyArray []Proxy

func (p *Proxy) GetAllUser() (err error, users ProxyArray) {
	resp, err := http.Get("http://127.0.0.1:9999/rest/api")
	if err != nil {
		return err, users
	}
	if resp.StatusCode != 200 {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, users
	}
	responseData := &common.Response{}
	err = json.Unmarshal(body, responseData)
	if err != nil {
		beeLogger.Log.Error("解析结果出错")
	}

	userArray := make([]byte, len(responseData.DataList))
	for i, data := range responseData.DataList {
		userArray[i] = data.(byte)
	}
	json.Unmarshal(userArray, users)
	return err, users
}

package gobaobi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var baseUrl = "https://api.baobi.com"

type Baobi struct {
	AccessKey, SecertKey, BaseUrl string
}

func NewBaobi(accessKey, secertKey, baseUrl string) *Baobi {
	return &Baobi{accessKey, secertKey, baseUrl}
}

func (this *Baobi) SetBaseUrl(url string) {
	baseUrl = url
	return
}

func (this *Baobi) signData(input map[string]interface{}) ([]byte, error) {
	input["nonce"] = strconv.FormatInt(time.Now().UnixNano(), 10)
	input["key"] = this.AccessKey
	qs := url.Values{}
	for k, v := range input {
		qs.Add(k, v.(string))
	}
	//md5SecertKey := fmt.Sprintf("%x", md5.Sum([]byte(this.SecertKey)))
	//h := hmac.New(sha256.New, []byte(md5SecertKey))
	h := hmac.New(sha256.New, []byte(this.SecertKey))
	io.WriteString(h, qs.Encode())
	signature := fmt.Sprintf("%x", h.Sum(nil))
	input["signature"] = signature
	requestData, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	return requestData, nil
}

func (this *Baobi) httpRequest(method string, path string, data []byte) ([]byte, error) {
	method = strings.ToUpper(method)
	api := baseUrl + path
	if this.BaseUrl != "" {
		api = this.BaseUrl + path
	}

	q := url.Values{}
	m := make(map[string]interface{})
	err := json.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	} else {
		for k, v := range m {
			q.Add(k, v.(string))
		}
	}

	req, err := http.NewRequest(method, api, strings.NewReader(q.Encode()))
	if err != nil {
		return nil, err
	}

	if "GET" == method {
		req.URL.RawQuery = q.Encode()
		//fmt.Println("===========================")
		//fmt.Println(req.URL.String())
	}

	if "POST" == method {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//fmt.Println("response URL:", resp.Request.URL)
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))
	//return ioutil.ReadAll(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func (this *Baobi) postRequest(path string, data map[string]interface{}) ([]byte, error) {
	requestData, _ := this.signData(data)
	body, err := this.httpRequest("post", path, requestData)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (this *Baobi) getRequest(path string, data map[string]interface{}) ([]byte, error) {
	requestData, _ := json.Marshal(data)
	body, err := this.httpRequest("get", path, requestData)
	if err != nil {
		return nil, err
	}
	return body, nil
}

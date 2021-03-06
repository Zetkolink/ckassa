package ckassa

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Shop struct {
	// Url к API Shop.
	Url string `json:"url"`

	// BaseUrl
	BaseUrl string `json:"baseUrl"`

	// Key уникальный ключ, используется для подписи сообщений.
	Key string `json:"key"`

	// Token идентификатор организации.
	Token string `json:"token"`

	// ServCode код провайдера.
	ServCode string `json:"servCode"`

	Cert *Certificate
}

func NewShop(url string, key string, token string, servCode string, certName string, certPath string, certPass string) (shop Shop, err error) {
	shop = Shop{
		Url:      url,
		BaseUrl:  url,
		Key:      key,
		Token:    token,
		ServCode: servCode,
	}

	shop.Cert, err = NewCert(certPath, certPass, certName)
	if err != nil {
		return
	}

	return
}

// SendRequest отправка запроса к Shop API.
func (s Shop) SendRequest(path string, data interface{}) ([]byte, *Response, error) {
	dataMap := GetStringMap(data)
	dataMap["sign"] = s.getSign(data)
	dataMap["shopToken"] = s.Token

	dataMapJson, _ := json.Marshal(dataMap)

	body := strings.NewReader(string(dataMapJson))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{s.Cert.Cert},
			},
		},
	}

	req, _ := http.NewRequest("POST", path, body)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("dn", s.Cert.Name)

	res, err := client.Do(req)
	if res != nil {
		defer func() {
			_ = res.Body.Close()
		}()
	}
	if err != nil {
		return nil, nil, err
	}
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusCreated {
		return contents, nil, nil
	}
	if string(contents) != "" {
		response, err := NewResponse(contents)
		if err != nil {
			return contents, nil, err
		}
		return contents, response, nil
	}

	return contents, nil, NotFound
}

// getSign создание подписи.
func (s Shop) getSign(data interface{}) string {
	return strings.ToUpper(GetMD5Hash(strings.ToUpper(GetMD5Hash(getSign(data) + s.Token + "&" + s.Key))))
}

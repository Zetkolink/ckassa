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

	// Key уникальный ключ, используется для подписи сообщений.
	Key string `json:"key"`

	// Token идентификатор организации.
	Token string `json:"token"`

	// ServCode код провайдера.
	ServCode string `json:"servCode"`

	Cert *Certificate
}

func NewShop(url string, key string, token string, servCode string, certName string, certPath string, certPass string) (shop *Shop, err error) {
	shop = &Shop{
		Url:      url + "/rs/shop",
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
func (s Shop) SendRequest(path string, data interface{}) (*Response, error) {
	dataMap := GetStringMap(data)
	dataMap["sign"] = s.getSign(dataMap)
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
		return nil, err
	}
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, NotFound
	}

	if res.StatusCode != http.StatusOK {
		return nil, ApiError
	}

	response, err := NewResponse(contents)
	if err != nil {
		return nil, ApiError
	}

	if string(contents) == "" {
		return nil, ApiError
	}

	return response, nil
}

// getSignString получение строки подписи из набора данных.
func (s Shop) getSignString(data map[string]string) string {
	values := GetValuesMap(data)
	return strings.Join(values, "&")
}

// getSign создание подписи.
func (s Shop) getSign(data map[string]string) string {
	return strings.ToUpper(GetMD5Hash(strings.ToUpper(GetMD5Hash(s.getSignString(data) + "&" + s.Token + "&" + s.Key))))
}

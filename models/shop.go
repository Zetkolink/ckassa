package models

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Shop struct {
	// Url к API Shop.
	Url string `json:"url"`

	// Key уникальный ключ, используется для подписи сообщений.
	Key string `json:"key"`

	// Token идентификатор организации.
	Token string `json:"token"`

	Cert *Certificate
}

func NewShop(url string, key string, token string, certPath string, certPass string) (shop *Shop, err error) {
	shop = &Shop{
		Url:   url + "/rs/shop",
		Key:   key,
		Token: token,
	}

	shop.Cert, err = NewCert(certPath, certPass)
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

	res, err := client.Post(path, "application/json", body)
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

	response, err := NewResponse(contents)
	if err != nil {
		return nil, ApiError
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("response with status code " + strconv.Itoa(res.StatusCode))
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

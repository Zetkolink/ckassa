package ckassa

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Shop struct {
	url   string
	key   string
	token string
	cert  *Certificate
}

func NewShop(url string, key string, token string, certPath string, certPass string) (shop *Shop, err error) {
	shop = &Shop{
		url:   url + "/rs/shop",
		key:   key,
		token: token,
	}

	shop.cert, err = NewCert(certPath, certPass)
	if err != nil {
		return
	}

	return
}

func (s Shop) SendRequest(path string, data interface{}) (*Response, error) {
	dataMap := getStringMap(data)
	dataMap["sign"] = s.getSign(dataMap)
	dataMap["shopToken"] = s.token

	dataMapJson, _ := json.Marshal(dataMap)

	body := strings.NewReader(string(dataMapJson))

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: s.cert.CertPool,
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
		return nil, apiError
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("response with status code " + strconv.Itoa(res.StatusCode))
	}

	if string(contents) == "" {
		return nil, apiError
	}

	return response, nil
}

func (s Shop) CreateMerchant(req MerchantRegRequest) (*Merchant, error) {
	path := s.url + RegMerchantPath
	resp, err := s.SendRequest(path, req)
	fmt.Println(resp)
	//merch := NewMerchant(resp.Body)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s Shop) getSignString(data map[string]string) string {
	values := getValuesMap(data)
	return strings.Join(values, "&")
}

func (s Shop) getSign(data map[string]string) string {
	return strings.ToUpper(getMD5Hash(strings.ToUpper(getMD5Hash(s.getSignString(data) + "&" + s.token + "&" + s.key))))
}

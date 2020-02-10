package ckassa

import "encoding/json"

// Payment платеж.
type Payment struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`

	// MethodType метод перехода по url для просмотра чека (GET/POST).
	MethodType string `json:"methodType"`

	// UserToken идентификатор пользователя
	UserToken string `json:"userToken"`

	// UserPhone номер телефона пользователя.
	UserPhone string `json:"userPhone"`

	// PayUrl url на страницу с чеком.
	PayUrl string `json:"payUrl"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`
}

func NewPayment(paymentJson []byte) (Payment, error) {
	pay := Payment{}

	err := json.Unmarshal(paymentJson, &pay)
	if err != nil {
		return pay, err
	}

	return pay, err
}

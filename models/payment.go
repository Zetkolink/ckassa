package models

import "encoding/json"

// Payment платеж.
type Payment struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"reg_pay_num"`

	// MethodType метод перехода по url для просмотра чека (GET/POST).
	MethodType string `json:"method_type"`

	// UserToken идентификатор пользователя
	UserToken string `json:"user_token"`

	// UserPhone номер телефона пользователя.
	UserPhone string `json:"user_phone"`

	// PayUrl url на страницу с чеком.
	PayUrl string `json:"pay_url"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchant_token"`
}

func NewPayment(paymentJson []byte) (Payment, error) {
	pay := Payment{}

	err := json.Unmarshal(paymentJson, &pay)
	if err != nil {
		return pay, err
	}
}

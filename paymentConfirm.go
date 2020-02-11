package ckassa

import "encoding/json"

// Payment платеж.
type PaymentConfirm struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`
}

func NewPaymentConfirm(paymentJson []byte) (PaymentConfirm, error) {
	pay := PaymentConfirm{}

	err := json.Unmarshal(paymentJson, &pay)
	if err != nil {
		return pay, err
	}

	return pay, err
}

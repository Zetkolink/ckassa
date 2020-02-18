package ckassa

import "encoding/json"

// Payment платеж.
type PaymentConfirm struct {
	// ResultState статус.
	ResultState string `json:"resultState"`

	// Desc описание.
	Desc string `json:"desc"`
}

func NewPaymentConfirm(paymentJson []byte) (PaymentConfirm, error) {
	pay := PaymentConfirm{}

	err := json.Unmarshal(paymentJson, &pay)
	if err != nil {
		return pay, err
	}

	return pay, err
}

package ckassa

import "encoding/json"

// Payment платеж.
type PaymentReservedUpdate struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`
}

func NewPaymentReservedUpdate(paymentJson []byte) (PaymentReservedUpdate, error) {
	pay := PaymentReservedUpdate{}

	err := json.Unmarshal(paymentJson, &pay)
	if err != nil {
		return pay, err
	}

	return pay, err
}

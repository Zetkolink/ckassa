package ckassa

import "encoding/json"

// Merchant мерчант.
type Merchant struct {
	// Phone номер телефона.
	Phone string `json:"phone"`

	// MerchantToken токен мерчанта.
	MerchantToken string `json:"merchantToken"`

	// State статус мерчанта.
	State string `json:"state"`
}

func NewMerchant(merchantJson []byte) (Merchant, error) {
	mch := Merchant{}
	err := json.Unmarshal(merchantJson, &mch)
	if err != nil {
		return mch, err
	}
	return mch, nil
}

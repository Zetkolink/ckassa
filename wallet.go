package ckassa

import (
	"encoding/json"
)

// Card карта.
type Wallet struct {
	// Balance баланс мерчанта в копейках.
	Balance int `json:"balance"`

	// WalletCode номер кошелька мерчанта в системе autopays.
	WalletCode string `json:"walletCode"`

	// MerchantToken токен мерчанта.
	MerchantToken string `json:"merchantToken"`
}

func NewWallet(walletJson []byte) (Wallet, error) {
	wallet := Wallet{}

	err := json.Unmarshal(walletJson, &wallet)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

package models

import (
	"encoding/json"
)

// CardReg регистрация карты.
type CardReg struct {
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
}

func NewCardReg(cardJson []byte) (CardReg, error) {
	cardReg := CardReg{}

	err := json.Unmarshal(cardJson, &cardReg)
	if err != nil {
		return cardReg, err
	}

	return cardReg, nil
}

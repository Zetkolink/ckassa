package ckassa

import (
	"encoding/json"
)

// Card карта.
type Card struct {
	// cardMask маскированный номер.
	CardMask string `json:"cardMask"`

	// exp время действия.
	Exp string `json:"exp"`

	// cardToken токен.
	CardToken string `json:"cardToken"`

	// cardType тип карты.
	CardType string `json:"cardType"`
}

func NewCard(cardJson []byte) (Card, error) {
	card := Card{}

	err := json.Unmarshal(cardJson, &card)
	if err != nil {
		return card, err
	}
	return card, nil
}

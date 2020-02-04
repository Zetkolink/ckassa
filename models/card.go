package models

import (
	"ckassa"
	"encoding/json"
)

// Card карта.
type Card struct {
	// cardMask маскированный номер.
	CardMask string `json:"card_mask"`

	// exp время действия.
	Exp string `json:"exp"`

	// cardToken токен.
	CardToken string `json:"card_token"`

	// cardType тип карты.
	CardType string `json:"card_type"`
}

func NewCard(cardJson []byte) (Card, error) {
	card := Card{}

	err := json.Unmarshal(cardJson, &card)
	if err != nil {
		return card, err
	}
	err = card.setCardType(card.CardType)
	if err != nil {
		return card, err
	}
	return card, nil
}

func (c Card) setCardType(cardType string) error {
	for _, v := range ckassa.CardTypes {
		if v == cardType {
			c.CardType = cardType
			return nil
		}
	}
	return ckassa.CardTypeNotAllowed
}

package ckassa

import "encoding/json"

// CardList список карт пользователя.
type CardList struct {
	// UserToken идентификатор пользователя
	UserToken string `json:"userToken"`

	// Cards список карт.
	Cards []Card `json:"cards"`
}

func NewCardList(cardListJson []byte) (CardList, error) {
	cardList := CardList{}

	err := json.Unmarshal(cardListJson, &cardList)
	if err != nil {
		return cardList, err
	}

	return cardList, nil
}

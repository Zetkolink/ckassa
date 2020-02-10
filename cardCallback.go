package ckassa

import "encoding/json"

// CardCallback уведомление о том, что карта зарегистрирована.
type CardCallback struct {
	// CardToken токен.
	CardToken string `json:"cardToken" gorm:"primary_key"`

	// UserToken идентификатор пользователя
	UserToken string `json:"userToken"`

	// CardMask маскированный номер.
	CardMask string `json:"cardMask"`

	// CardType тип карты.
	CardType string `json:"cardType"`

	// ClientInfo идентификатор клиента.
	ClientInfo string `json:"clientInfo"`

	// ShopToken время действия.
	ShopToken string `json:"shopToken"`

	// Sign подпись.
	Sign string `json:"sign"`
}

func NewCardCallback(cardCallbackJson []byte) (CardCallback, error) {
	cardCallback := CardCallback{}

	err := json.Unmarshal(cardCallbackJson, &cardCallback)
	if err != nil {
		return cardCallback, err
	}

	return cardCallback, nil
}

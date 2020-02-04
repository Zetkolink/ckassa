package ckassa

type Card struct {
	// cardMask маскированный номер.
	cardMask string

	// exp время действия.
	exp string

	// cardToken токен.
	cardToken string

	// cardType тип карты.
	cardType string
}

func NewCard(cardMask string, exp string, cardToken string, cardType string) (*Card, error) {
	card := &Card{
		cardMask:  cardMask,
		exp:       exp,
		cardToken: cardToken,
	}

	err := card.setCardType(cardType)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (c Card) setCardMask(cardMask string) {
	c.cardMask = cardMask
}

func (c Card) setExp(exp string) {
	c.exp = exp
}

func (c Card) setCardToken(cardToken string) {
	c.cardToken = cardToken
}

func (c Card) setCardType(cardType string) error {
	for _, v := range CardTypes {
		if v == cardType {
			c.cardType = cardType
			return nil
		}
	}
	return cardTypeNotAllowed
}

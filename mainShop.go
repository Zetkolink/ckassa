package ckassa

// MerchantShop структура для работы с Merchant Shop API.
type MainShop struct {
	*Shop
}

func NewMainShop(sh *Shop) *MainShop {
	return &MainShop{sh}
}

// CreateMerchant регистрация мерчанта.
func (m MainShop) CreateUser(req UserRegRequest) (*User, error) {
	path := m.Url + UserRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	user, err := NewUser([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// LoadUser получение данных зарегестрированного пользователя.
func (m MainShop) LoadUser(req MerchantRegRequest) (*User, error) {
	path := m.Url + UserStatusPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	user, err := NewUser([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) RegCard(req CardRegRequest) (*CardReg, error) {
	path := m.Url + CardRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	cardReg, err := NewCardReg([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &cardReg, nil
}

// GetCardList получение списка карт.
func (m MainShop) GetCardList(req CardRegRequest) (*CardList, error) {
	path := m.Url + GetCardListPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	cardList, err := NewCardList([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &cardList, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) DeactivateCard(req CardDeactivateRequest) (*Result, error) {
	path := m.Url + CardDeactivationPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	result, err := NewResult([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

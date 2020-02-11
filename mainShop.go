package ckassa

// MerchantShop структура для работы с Merchant Shop API.
type MainShop struct {
	*Shop
}

func NewMainShop(sh *Shop) *MainShop {
	return &MainShop{sh}
}

// CreateMerchant регистрация мерчанта.
func (m MainShop) CreateUser(req UserRegRequest) (*User, *Response, error) {
	path := m.Url + UserRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	user, err := NewUser([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &user, resp, nil
}

// LoadUser получение данных зарегестрированного пользователя.
func (m MainShop) LoadUser(req MerchantRegRequest) (*User, *Response, error) {
	path := m.Url + UserStatusPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	user, err := NewUser([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &user, resp, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) RegCard(req CardRegRequest) (*CardReg, *Response, error) {
	path := m.Url + CardRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	cardReg, err := NewCardReg([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &cardReg, resp, nil
}

// GetCardList получение списка карт.
func (m MainShop) GetCardList(req CardRegRequest) (*CardList, *Response, error) {
	path := m.Url + GetCardListPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	cardList, err := NewCardList([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &cardList, resp, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) DeactivateCard(req CardDeactivateRequest) (*Result, *Response, error) {
	path := m.Url + CardDeactivationPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	result, err := NewResult([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &result, resp, nil
}

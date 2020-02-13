package ckassa

// MerchantShop структура для работы с Merchant Shop API.
type MainShop struct {
	*Shop
}

func NewMainShop(sh *Shop) *MainShop {
	ms := &MainShop{sh}
	ms.Url += "/rs/shop"
	return ms
}

// CreateMerchant регистрация мерчанта.
func (m MainShop) CreateUser(req UserRegRequest) (*User, *Response, error) {
	path := m.Url + UserRegPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}

	user, err := NewUser(resp)
	if err != nil {
		return nil, nil, err
	}
	return &user, nil, nil
}

// LoadUser получение данных зарегестрированного пользователя.
func (m MainShop) LoadUser(req LoadUserRequest) (*User, *Response, error) {
	path := m.Url + UserStatusPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	user, err := NewUser(resp)
	if err != nil {
		return nil, nil, err
	}

	return &user, nil, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) RegCard(req CardRegRequest) (*CardReg, *Response, error) {
	path := m.Url + CardRegPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	cardReg, err := NewCardReg(resp)
	if err != nil {
		return nil, nil, err
	}

	return &cardReg, nil, nil
}

// GetCardList получение списка карт.
func (m MainShop) GetCardList(req CardRegRequest) (*CardList, *Response, error) {
	path := m.Url + GetCardListPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	cardList, err := NewCardList(resp)
	if err != nil {
		return nil, nil, err
	}

	return &cardList, nil, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) DeactivateCard(req CardDeactivateRequest) (*Result, *Response, error) {
	path := m.Url + CardDeactivationPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	result, err := NewResult(resp)
	if err != nil {
		return nil, nil, err
	}

	return &result, nil, nil
}

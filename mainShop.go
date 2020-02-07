package ckassa

import "ckassa/models"

// MerchantShop структура для работы с Merchant Shop API.
type MainShop struct {
	models.Shop
}

func NewMainShop(sh models.Shop) *MainShop {
	return &MainShop{sh}
}

// CreateMerchant регистрация мерчанта.
func (m MainShop) CreateUser(req MerchantRegRequest) (*models.User, error) {
	path := m.Url + UserRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	user, err := models.NewUser([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// LoadUser получение данных зарегестрированного пользователя.
func (m MainShop) LoadUser(req MerchantRegRequest) (*models.User, error) {
	path := m.Url + UserStatusPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	user, err := models.NewUser([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) RegCard(req CardRegRequest) (*models.CardReg, error) {
	path := m.Url + CardRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	cardReg, err := models.NewCardReg([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &cardReg, nil
}

// GetCardList получение списка карт.
func (m MainShop) GetCardList(req CardRegRequest) (*models.CardList, error) {
	path := m.Url + GetCardListPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	cardList, err := models.NewCardList([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &cardList, nil
}

// DeactivateCard деактивация карты.
func (m MainShop) DeactivateCard(req CardDeactivateRequest) (*models.Result, error) {
	path := m.Url + CardDeactivationPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	result, err := models.NewResult([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &result, nil
}

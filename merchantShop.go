package ckassa

import "ckassa/models"

// MerchantShop структура для работы с Merchant Shop API.
type MerchantShop struct {
	models.Shop
}

func NewMerchantShop(sh models.Shop) *MerchantShop {
	return &MerchantShop{sh}
}

// CreateMerchant регистрация мерчанта.
func (m MerchantShop) CreateMerchant(req MerchantRegRequest) (*models.Merchant, error) {
	path := m.Url + MerchantRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	merchant, err := models.NewMerchant([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &merchant, nil
}

// CreateMobilePayment создание платежа в пользу мерчанта, оплата с баланса сотового телефона.
func (m MerchantShop) CreateMobilePayment(req MobilePaymentRequest) (*models.Payment, error) {
	path := m.Url + MobilePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	payment, err := models.NewPayment([]byte(resp.Body))
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// CreatePayment создание платежа в пользу мерчанта.
func (m MerchantShop) CreatePayment(req PaymentRequest) (*models.Payment, error) {
	path := m.Url + PaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	payment, err := models.NewPayment([]byte(resp.Body))
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// GetBalance получение баланса мерчанта.
func (m MerchantShop) GetBalance(req GetBalanceRequest) (*Response, error) {
	path := m.Url + BalancePath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// LoadMerchant получение данных зарегистрированного мерчанта.
func (m MerchantShop) LoadMerchant(req LoadMerchantRequest) (*models.Merchant, error) {
	path := m.Url + MerchantStatusPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	merchant, err := models.NewMerchant([]byte(resp.Body))
	if err != nil {
		return nil, err
	}

	return &merchant, nil
}

// ReservePayment бронирование средств на балансе пользователя.
func (m MerchantShop) ReservePayment(req PaymentRequest) (*models.Payment, error) {
	path := m.Url + ReservePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	payment, err := models.NewPayment([]byte(resp.Body))
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (m MerchantShop) ConfirmReservePayment(req UpdatePayMerchantRequest) (*models.Payment, error) {
	path := m.Url + ConfirmReservePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	payment, err := models.NewPayment([]byte(resp.Body))
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

// UpdatePayMerchant обновление получателя ( мерчанта ) платежа.
func (m MerchantShop) UpdatePayMerchant(req UpdatePayMerchantRequest) (*models.Payment, error) {
	path := m.Url + UpdatePayMerchantPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, err
	}
	payment, err := models.NewPayment([]byte(resp.Body))
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

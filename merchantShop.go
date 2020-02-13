package ckassa

// MerchantShop структура для работы с Merchant Shop API.
type MerchantShop struct {
	Shop
}

func NewMerchantShop(sh Shop) *MerchantShop {
	ms := &MerchantShop{sh}
	ms.Url += "/rs/merchant"
	return ms
}

// CreateMerchant регистрация мерчанта.
func (m MerchantShop) CreateMerchant(req MerchantRegRequest) (*Merchant, *Response, error) {
	path := m.Url + MerchantRegPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	merchant, err := NewMerchant(resp)
	if err != nil {
		return nil, nil, err
	}

	return &merchant, nil, nil
}

// CreateMobilePayment создание платежа в пользу мерчанта, оплата с баланса сотового телефона.
func (m MerchantShop) CreateMobilePayment(req MobilePaymentRequest) (*Payment, *Response, error) {
	path := m.Url + MobilePaymentPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	payment, err := NewPayment(resp)
	if err != nil {
		return nil, nil, err
	}
	return &payment, nil, nil
}

// CreatePayment создание платежа в пользу мерчанта.
func (m MerchantShop) CreatePayment(req PaymentRequest) (*Payment, *Response, error) {
	path := m.Url + PaymentPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	payment, err := NewPayment(resp)
	if err != nil {
		return nil, nil, err
	}
	return &payment, nil, nil
}

// GetBalance получение баланса мерчанта.
func (m MerchantShop) GetBalance(req GetBalanceRequest) (*Wallet, *Response, error) {
	path := m.Url + BalancePath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	wallet, err := NewWallet(resp)
	if err != nil {
		return nil, nil, err
	}
	return &wallet, nil, nil
}

// LoadMerchant получение данных зарегистрированного мерчанта.
func (m MerchantShop) LoadMerchant(req LoadMerchantRequest) (*Merchant, *Response, error) {
	path := m.Url + MerchantStatusPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	merchant, err := NewMerchant(resp)
	if err != nil {
		return nil, nil, err
	}

	return &merchant, nil, nil
}

// ReservePayment бронирование средств на балансе пользователя.
func (m MerchantShop) ReservePayment(req PaymentRequest) (*Payment, *Response, error) {
	path := m.Url + ReservePaymentPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	payment, err := NewPayment(resp)
	if err != nil {
		return nil, nil, err
	}
	return &payment, nil, nil
}

// ConfirmReservePayment подтверждение забронированного платежа.
func (m MerchantShop) ConfirmReservePayment(req UpdatePayMerchantRequest) (*PaymentReservedUpdate, *Response, error) {
	path := m.Url + ConfirmReservePaymentPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	payment, err := NewPaymentReservedUpdate(resp)
	if err != nil {
		return nil, nil, err
	}
	return &payment, nil, nil
}

// UpdatePayMerchant обновление получателя ( мерчанта ) платежа.
func (m MerchantShop) UpdatePayMerchant(req UpdatePayMerchantRequest) (*PaymentReservedUpdate, *Response, error) {
	path := m.Url + UpdatePayMerchantPath
	resp, errResp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if errResp != nil {
		return nil, errResp, nil
	}
	payment, err := NewPaymentReservedUpdate(resp)
	if err != nil {
		return nil, nil, err
	}
	return &payment, nil, nil
}

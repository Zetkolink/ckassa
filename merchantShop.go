package ckassa

// MerchantShop структура для работы с Merchant Shop API.
type MerchantShop struct {
	*Shop
}

func NewMerchantShop(sh *Shop) *MerchantShop {
	return &MerchantShop{sh}
}

// CreateMerchant регистрация мерчанта.
func (m MerchantShop) CreateMerchant(req MerchantRegRequest) (*Merchant, *Response, error) {
	path := m.Url + MerchantRegPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	merchant, err := NewMerchant([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &merchant, resp, nil
}

// CreateMobilePayment создание платежа в пользу мерчанта, оплата с баланса сотового телефона.
func (m MerchantShop) CreateMobilePayment(req MobilePaymentRequest) (*Payment, *Response, error) {
	path := m.Url + MobilePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	payment, err := NewPayment([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}
	return &payment, resp, nil
}

// CreatePayment создание платежа в пользу мерчанта.
func (m MerchantShop) CreatePayment(req PaymentRequest) (*Payment, *Response, error) {
	path := m.Url + PaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	payment, err := NewPayment([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}
	return &payment, resp, nil
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
func (m MerchantShop) LoadMerchant(req LoadMerchantRequest) (*Merchant, *Response, error) {
	path := m.Url + MerchantStatusPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	merchant, err := NewMerchant([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}

	return &merchant, resp, nil
}

// ReservePayment бронирование средств на балансе пользователя.
func (m MerchantShop) ReservePayment(req PaymentRequest) (*Payment, *Response, error) {
	path := m.Url + ReservePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	payment, err := NewPayment([]byte(*resp.Body))
	if err != nil {
		return nil, resp, err
	}
	return &payment, resp, nil
}

// ConfirmReservePayment подтверждение забронированного платежа.
func (m MerchantShop) ConfirmReservePayment(req UpdatePayMerchantRequest) (*PaymentReservedUpdate, *Response, error) {
	path := m.Url + ConfirmReservePaymentPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	payment, err := NewPaymentReservedUpdate([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}
	return &payment, resp, nil
}

// UpdatePayMerchant обновление получателя ( мерчанта ) платежа.
func (m MerchantShop) UpdatePayMerchant(req UpdatePayMerchantRequest) (*PaymentReservedUpdate, *Response, error) {
	path := m.Url + UpdatePayMerchantPath
	resp, err := m.SendRequest(path, req)
	if err != nil {
		return nil, nil, err
	}
	if resp.Body == nil {
		return nil, resp, nil
	}
	payment, err := NewPaymentReservedUpdate([]byte(*resp.Body))
	if err != nil {
		return nil, nil, err
	}
	return &payment, resp, nil
}

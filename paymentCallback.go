package ckassa

import "encoding/json"

// PaymentCallback уведомление о статусе платежа.
type PaymentCallback struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum" gorm:"primary_key"`

	// State статус платежа
	State string `json:"state"`

	// TotalAmount сумма платежа в копейках.
	TotalAmount int `json:"totalAmount"`

	// CreatedDate время создания платежа (Время должно быть в формате: yyyy-mm-dd hh:mm:ss).
	CreatedDate string `json:"createdDate"`

	// ProviderServCode уникальный код провайдера услуг.
	ProviderServCode string `json:"providerServCode"`

	// ProviderName название провайдера услуг.
	ProviderName string `json:"providerName"`

	// ErrorCode код ошибки, используется для локализации проблемы.
	ErrorCode string `json:"errorCode"`

	// Error текст ошибки.
	Error string `json:"error"`

	// Message сообщение пользователю.
	Message string `json:"message"`

	// ProvisionServices флаг, была ли оказана услуга.
	ProvisionServices bool `json:"provisionServices"`

	// ProcDate время проведения платежа в формате yyyy-mm-dd hh:mm:ss.
	ProcDate string `json:"procDate"`

	// Sign подпись.
	Sign string `json:"sign"`
}

func NewPaymentCallback(paymentCallbackJson []byte) (PaymentCallback, error) {
	paymentCallback := PaymentCallback{}

	err := json.Unmarshal(paymentCallbackJson, &paymentCallback)
	if err != nil {
		return paymentCallback, err
	}

	return paymentCallback, nil
}

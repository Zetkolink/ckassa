package ckassa

type MerchantRegRequest struct {
	// Phone номер телефона в формате 7##########.
	Phone string `json:"phone" validate:"required"`

	// Email электронная почта.
	Email *string `json:"email"`

	// Name имя мерчатнта.
	Name string `json:"name" validate:"required"`

	// SurName фамилия.
	SurName string `json:"surName" validate:"required"`

	// MiddleName отчество.
	MiddleName string `json:"middleName"`

	// CallName позывной.
	CallName string `json:"callName" validate:"required"`

	// Region регион, либо город мерчанта.
	Region *string `json:"region"`

	// docList список документов (используется для идентификации).
	DocList *Docs
}

// Doc информация о документе мерчанта.
type Doc struct {
	// Type тип документа
	// passport - паспорт
	// snils - СНИЛС
	// inn - ИНН
	// car_driver - водительское удостоверение.
	Type string `json:"type"`

	// Number номер документа
	Number string `json:"number"`
}

// Docs список документов мерчанта.
type Docs []Doc

// MobilePaymentRequest тело запроса на создание платежа в пользу мерчанта с баланса мобильного телефона.
type MobilePaymentRequest struct {
	// ServiceCode код провайдера.
	ServiceCode string `json:"serviceCode"`

	// Amount сумма платежа в копейках, которая идет на счет пользователю.
	Amount string `json:"amount"`

	// Comission комиссия платежа в копейках, при отсутствии комиссии передается 0.
	Comission string `json:"comission"`

	// OrderId уникальный идентификатор заказа на стороне магазина.
	OrderId string `json:"orderId"`

	// Description дополнительные сведения.
	Description string `json:"description"`

	// UserToken идентификатор пользователя, если клиент был ранее зарегестрирован
	// то оплата будет произведена с телефона указанного при регистрации (login).
	UserToken string `json:"userToken"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`

	// UserPhone номер телефона для СМС платежа с которого будет произведена оплата,
	// должен отправляться только в случае если пользователь не зарегестрирован.
	UserPhone string `json:"userPhone"`

	// UserEmail e-mail для отправки фискального чека.
	UserEmail string `json:"userEmail"`

	// FiscalType нужно ли отправлять фискальный чек.
	FiscalType string `json:"fiscalType"`
}

// PaymentRequest тело запроса на создание платежа в пользу мерчанта.
type PaymentRequest struct {
	// ServiceCode код провайдера.
	ServiceCode string `json:"serviceCode"`

	// Amount сумма платежа в копейках, которая идет на счет пользователю.
	Amount string `json:"amount"`

	// Comission комиссия платежа в копейках, при отсутствии комиссии передается 0.
	Comission string `json:"comission"`

	// OrderId уникальный идентификатор заказа на стороне магазина.
	OrderId string `json:"orderId"`

	// Description дополнительные сведения.
	Description string `json:"description"`

	// UserToken идентификатор пользователя, если клиент был ранее зарегестрирован
	// то оплата будет произведена с телефона указанного при регистрации (login).
	UserToken string `json:"userToken"`

	// CardToken идентификатор карты.
	CardToken string `json:"cardToken"`

	// GPayToken токен при оплате через Google Pay.
	GPayToken string `json:"gPayToken"`

	// EnableSMSConfirm Используется только для рекуррентных платежей
	// Данный параметр определяет включать ли поддержку 3ds.
	EnableSMSConfirm string `json:"enableSMSConfirm"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`

	// CallName позывной мерчанта, используется для дополнительной идентификации.
	CallName string `json:"callName"`

	// ExtraPhone дополнительный телефон мерчанта (используется для дополнительной идентификации).
	ExtraPhone string `json:"extraPhone"`

	// HoldTtl время заморозки средств (в секундах). Минимум 10 минут, максимум 4 дня, по умолчанию 30 минут,
	// если услуга не будет подтверждена, произойдет автоматический возврат средств на карту клиента.
	HoldTtl string `json:"holdTtl"`

	// PayType тип оплаты.
	PayType string `json:"payType"`

	// FiscalType нужно ли отправлять фискальный чек.
	FiscalType string `json:"fiscalType"`

	// UserEmail email пользователя для отправки чека.
	UserEmail string `json:"userEmail"`
}

// GetPaymentStateRequest тело запроса на получение статуса платежа.
type GetPaymentStateRequest struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`
}

// GetBalanceRequest тело запроса на получение баланса мерчанта.
type GetBalanceRequest struct {
	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`
}

// LoadMerchantRequest тело запроса на получение данных мерчанта.
type LoadMerchantRequest struct {
	// Login номер телефона в формате 7**********.
	Login string `json:"login"`
}

// UpdatePayMerchantRequest тело запроса на обновление получателя ( мерчанта ) платежа.
type UpdatePayMerchantRequest struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchantToken"`

	// CallName позывной мерчанта, используется для дополнительной идентификации.
	CallName string `json:"callName"`

	// ExtraPhone дополнительный телефон мерчанта (используется для дополнительной идентификации).
	ExtraPhone string `json:"extraPhone"`
}

// ConfirmPaymentRequest тело запроса на подтверждение оказания услуги.
type ConfirmPaymentRequest struct {
	// RegPayNum номер платежа.
	RegPayNum string `json:"regPayNum"`

	// OrderId уникальный идентификатор заказа на стороне магазина.
	OrderId string `json:"orderId"`

	// Amount сумма платежа в копейках, которая идет на счет пользователю.
	Amount string `json:"amount"`
}

// UserRegRequest тело запроса на регистрацию пользователя.
type UserRegRequest struct {
	// Login номер телефона в формате 7**********.
	Login string `json:"login" validate:"required"`

	// Email почта.
	Email *string `json:"email"`

	// Name имя.
	Name *string `json:"name"`

	// Surname фамилия.
	SurName *string `json:"surName"`

	// MiddleName отчество.
	MiddleName *string `json:"middleName"`
}

// LoadUserRequest тело запроса на получение данных пользователя.
type LoadUserRequest struct {
	// Login номер телефона в формате 7**********.
	Login string `json:"login"`
}

// CardRegRequest тело запроса на регистрацию карты.
type CardRegRequest struct {
	// ClientType устанавливает вид формы оплаты (та, что будет отправлена клиенту)
	// * web - WEB форма (по умолчанию)
	// * mobile - мобильная форма
	// * iframe - фрейм
	ClientType string `json:"clientType"`

	// UserToken идентификатор пользователя.
	UserToken string `json:"userToken"`
}

// GetCardListRequest тело запроса на получения списка карт.
type GetCardListRequest struct {
	// UserToken идентификатор пользователя.
	UserToken string `json:"userToken"`
}

// CardDeactivateRequest тело запроса на деактивацию зарегитстрированной карты.
type CardDeactivateRequest struct {
	// UserToken идентификатор пользователя.
	UserToken string `json:"userToken"`

	// CardToken идентификатор карты.
	CardToken string `json:"cardToken"`
}

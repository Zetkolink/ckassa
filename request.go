package ckassa

type MerchantRegRequest struct {
	// Phone номер телефона в формате 7##########.
	Phone string `json:"phone"`

	// Email электронная почта.
	Email string `json:"email"`

	// Name имя мерчатнта.
	Name string `json:"name"`

	// SurName фамилия.
	SurName string `json:"surName"`

	// MiddleName отчество.
	MiddleName string `json:"middleName"`

	// CallName позывной.
	CallName string `json:"callName"`

	// Region регион, либо город мерчанта.
	Region string `json:"region"`

	// docList список документов (используется для идентификации).
	DocList Docs
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

type MobilePaymentRequest struct {
	// ServiceCode код провайдера.
	ServiceCode string `json:"service_code"`

	// Amount сумма платежа в копейках, которая идет на счет пользователю.
	Amount string `json:"amount"`

	// Comission комиссия платежа в копейках, при отсутствии комиссии передается 0.
	Comission string `json:"comission"`

	// OrderId уникальный идентификатор заказа на стороне магазина.
	OrderId string `json:"order_id"`

	// Description дополнительные сведения.
	Description string `json:"description"`

	// UserToken идентификатор пользователя, если клиент был ранее зарегестрирован
	// то оплата будет произведена с телефона указанного при регистрации (login).
	UserToken string `json:"user_token"`

	// MerchantToken идентификатор мерчанта.
	MerchantToken string `json:"merchant_token"`

	// UserPhone номер телефона для СМС платежа с которого будет произведена оплата,
	// должен отправляться только в случае если пользователь не зарегестрирован.
	UserPhone string `json:"user_phone"`

	// UserEmail e-mail для отправки фискального чека.
	UserEmail string `json:"user_email"`

	// FiscalType нужно ли отправлять фискальный чек.
	FiscalType string `json:"fiscal_type"`
}

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

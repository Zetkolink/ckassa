package ckassa

import "errors"

var (
	// CertificateNotFound файл сертификата не найден.
	CertificateNotFound = errors.New("certificate file not found")

	// CardTypeNotAllowed данный тип карт не поддерживается.
	CardTypeNotAllowed = errors.New("card type not allowed")

	// UserStateNotAllowed данный статус пользователя не поддерживается.
	UserStateNotAllowed = errors.New("user state not allowed")

	// ApiError api вернул неправильный ответ.
	ApiError = errors.New("api return bad body")
)

package ckassa

import "errors"

var (
	// certificateNotFound файл сертификата не найден.
	certificateNotFound = errors.New("certificate file not found")

	// cardTypeNotAllowed данный тип карт не поддерживается.
	cardTypeNotAllowed = errors.New("card type not allowed")

	// userStateNotAllowed данный статус пользователя не поддерживается.
	userStateNotAllowed = errors.New("user state not allowed")

	// apiError api вернул неправильный ответ.
	apiError = errors.New("api return bad body")
)

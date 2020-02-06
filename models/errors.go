package models

import "errors"

// ApiError api вернул неправильный ответ.
var ApiError = errors.New("api return bad body")

// CertificateNotFound файл сертификата не найден.
var CertificateNotFound = errors.New("certificate file not found")

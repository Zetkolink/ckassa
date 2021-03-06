package ckassa

import "errors"

// ApiError api вернул неправильный ответ.
var ApiError = errors.New("api return bad body")

// NotFound не найдено.
var NotFound = errors.New("item not found")

// CertificateNotFound файл сертификата не найден.
var CertificateNotFound = errors.New("certificate file not found")

// CertificateDecryptError ошибка при расшифровке сертификата.
var CertificateDecryptError = errors.New("certificate decrypt error")

// CertificatePairError ошибка при создании сертификата.
var CertificatePairError = errors.New("certificate pair error")

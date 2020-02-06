package models

import (
	"crypto/x509"
	"io/ioutil"
)

// Cert сертификат.
type Certificate struct {
	// Path путь до файла сертификата.
	Path string `json:"path"`

	// Pass пароль.
	Pass string `json:"pass"`

	// Name имя сертификата.
	Name string `json:"name"`

	CertPool *x509.CertPool
}

// NewCert загрузка сертификата.
func NewCert(path string, pass string) (*Certificate, error) {
	cert := &Certificate{
		Path: path,
		Pass: pass,
	}

	if !fileExists(path) {
		return cert, CertificateNotFound
	}

	caCert, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert.CertPool = caCertPool

	return cert, nil
}

package models

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
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

	Cert tls.Certificate
}

// NewCert загрузка сертификата.
func NewCert(path string, pass string, name string) (*Certificate, error) {
	cert := &Certificate{
		Path: path,
		Pass: pass,
		Name: name,
	}

	if !fileExists(path) {
		return cert, CertificateNotFound
	}

	bundle, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	keyBlock, certsPEM := pem.Decode(bundle)

	keyDER, err := x509.DecryptPEMBlock(keyBlock, []byte(pass))
	if err != nil {
		return nil, CertificateDecryptError
	}

	keyBlock.Bytes = keyDER
	keyBlock.Headers = nil

	keyPEM := pem.EncodeToMemory(keyBlock)

	certificate, err := tls.X509KeyPair(certsPEM, keyPEM)
	if err != nil {
		return nil, CertificatePairError
	}

	cert.Cert = certificate

	return cert, nil
}

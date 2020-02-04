package ckassa

import (
	"crypto/md5"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"os"
	"reflect"
	"strconv"
)

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// decodeX509 расшифровка сертификата.
func decodeX509(pubPem []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(pubPem)

	pub, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}

	return pub, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getStringMap(data interface{}) map[string]string {
	sVal := reflect.ValueOf(data)
	sType := reflect.TypeOf(data)
	length := sVal.NumField()
	values := make(map[string]string, 0)
	for i := 0; i < length; i++ {
		field := sVal.Field(i)
		key := sType.Field(i).Name
		value := ""
		switch field.Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int:
			value = strconv.Itoa(int(field.Int()))
		}

		if value != "" {
			values[key] = value
		}

	}
	return values
}

func getKeysMap(mp map[string]string) []string {
	keys := make([]string, len(mp))

	for k, _ := range mp {
		keys = append(keys, k)
	}

	return keys
}

func getValuesMap(mp map[string]string) []string {
	values := make([]string, len(mp))

	for _, v := range mp {
		values = append(values, v)
	}

	return values
}

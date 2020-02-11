package ckassa

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"reflect"
	"strconv"
)

func GetStringMap(data interface{}) map[string]string {
	sVal := reflect.ValueOf(data)
	sType := reflect.TypeOf(data)
	length := sVal.NumField()
	values := make(map[string]string, 0)
	for i := 0; i < length; i++ {
		field := sVal.Field(i)
		key := sType.Field(i).Tag.Get("json")
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

func GetMD5Hash(text string) string {
	hs := md5.New()
	hs.Write([]byte(text))
	return hex.EncodeToString(hs.Sum(nil))
}

func GetValuesMap(mp map[string]string) []string {
	values := make([]string, 0)
	for _, v := range mp {
		values = append(values, v)
	}
	return values
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

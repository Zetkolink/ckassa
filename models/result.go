package models

import (
	"encoding/json"
)

// Result результат.
type Result struct {
	// ResultState результат
	// * success - успех
	// * fail - неудача
	// * desc - описание
	ResultState string `json:"resultState"`

	// UserToken идентификатор пользователя
	UserToken string `json:"userToken"`
}

func NewResult(resultJson []byte) (Result, error) {
	result := Result{}

	err := json.Unmarshal(resultJson, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

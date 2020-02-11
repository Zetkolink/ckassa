package ckassa

import (
	"encoding/json"
)

// User пользователь.
type User struct {
	// Login номер телефона в формате 7**********.
	Login string `json:"login"`

	// Email почта.
	Email string `json:"email"`

	// Name имя.
	Name string `json:"name"`

	// Surname фамилия.
	Surname string `json:"surName"`

	// MiddleName отчество.
	MiddleName string `json:"middleName"`

	// State статус.
	State string `json:"state"`

	// UserToken идентификатор пользователя.
	UserToken string `json:"userToken"`
}

func NewUser(userJson []byte) (User, error) {
	user := User{}

	err := json.Unmarshal(userJson, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

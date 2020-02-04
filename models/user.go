package models

import (
	"ckassa"
	"encoding/json"
)

// User пользователь.
type User struct {
	// Login номер телефона.
	Login string `json:"login"`

	// Email почта.
	Email string `json:"email"`

	// Name имя.
	Name string `json:"name"`

	// Surname фамилия.
	Surname string `json:"surname"`

	// MiddleName отчество.
	MiddleName string `json:"middle_name"`

	// State статус.
	State string `json:"state"`

	// UserToken идентификатор пользователя.
	UserToken string `json:"user_token"`
}

func NewUser(userJson []byte) (User, error) {
	user := User{}

	err := json.Unmarshal(userJson, &user)
	if err != nil {
		return user, err
	}

	err = user.setState(user.State)
	if err != nil {
		return user, err
	}

	return user, nil
}

// setState изменение статуса пользователя.
func (u User) setState(state string) error {
	for _, v := range ckassa.UserStates {
		if v == state {
			u.State = state
			return nil
		}
	}
	return ckassa.UserStateNotAllowed
}

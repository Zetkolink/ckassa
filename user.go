package ckassa

// User пользователь.
type User struct {
	// login номер телефона.
	login string

	// email почта.
	email string

	// name имя.
	name string

	// surname фамилия.
	surname string

	// middleName отчество.
	middleName string

	// state статус.
	state string

	// userToken идентификатор пользователя.
	userToken string
}

func NewUser(login string, email string, name string, sName string, mName string, state string, userToken string) (User, error) {
	user := User{
		login:      login,
		email:      email,
		name:       name,
		surname:    sName,
		middleName: mName,
		userToken:  userToken,
	}

	err := user.setState(state)
	if err != nil {
		return user, err
	}

	return user, nil
}

// setState изменение статуса пользователя.
func (u User) setState(state string) error {
	for _, v := range UserStates {
		if v == state {
			u.state = state
			return nil
		}
	}
	return userStateNotAllowed
}

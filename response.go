package ckassa

import "encoding/json"

// Response струтура для работы с ответами ShopAPI.
type Response struct {
	// Code код ответа.
	Code string `json:"code"`

	// Message сообщение в ответе.
	Message string `json:"message"`

	// UserMessage сообщение для пользователя.
	UserMessage string `json:"user_message"`

	// Body тело ответа.
	Body string `json:"body"`
}

func NewResponse(responseBytes []byte) (*Response, error) {
	resp := &Response{}

	err := json.Unmarshal(responseBytes, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

package ckassa

import "encoding/json"

// Response струтура для работы с ответами ShopAPI.
type Response struct {
	// Sign подпись.
	Sign string `json:"sign"`

	// Code код ответа.
	Code int `json:"code"`

	// Message сообщение в ответе.
	Message string `json:"message"`

	// UserMessage сообщение для пользователя.
	UserMessage string `json:"userMessage"`

	// Body тело ответа.
	Body *string `json:"body"`
}

func NewResponse(responseBytes []byte) (*Response, error) {
	resp := &Response{}

	err := json.Unmarshal(responseBytes, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

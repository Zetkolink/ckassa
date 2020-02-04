package ckassa

type Payment struct {
	regPayNum     string
	methodType    string
	userToken     string
	userPhone     string
	payUrl        string
	merchantToken string
}

func NewPayment(regPayNum string, methodType string, userToken string, userPhone string, payUrl string, merchantToken string) *Payment {
	return &Payment{
		regPayNum:     regPayNum,
		methodType:    methodType,
		userToken:     userToken,
		userPhone:     userPhone,
		payUrl:        payUrl,
		merchantToken: merchantToken,
	}
}

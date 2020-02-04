package ckassa

type Merchant struct {
	phone         string
	merchantToken string
	state         string
}

func NewMerchant(phone string, state string, merchantToken string) *Merchant {
	return &Merchant{
		phone:         phone,
		merchantToken: merchantToken,
		state:         state,
	}
}

package ckassa

import "log"

func main() {
	shp, err := NewShop(
		"534a63d0-2acb-4273-b4d4-25536f578483",
		"bb9f850d-86f7-4bd7-957c-282b79037c5d",
		"/C=RU/ST=Permskiy kray/L=Perm/O=OOO Billing Systems/OU=BS_CA/CN=59-BS-TERM-Test_ShopApi-001",
		"e9shpkwgyc",
	)
	if err != nil {
		log.Fatal(err)
	}

	regM := MerchantRegRequest{
		Phone:      "89044526201",
		Email:      "z.chalimov@yandex.ru",
		Name:       "Zakir",
		SurName:    "Chalimov",
		MiddleName: "Zagidovich",
		CallName:   "zetkolink",
		Region:     "Tyumen",
		DocList:    Docs{
			Doc{
				Type:   "passport",
				Number: "499123",
			},
		},
	}

	_, err = shp.CreateMerchant(regM)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

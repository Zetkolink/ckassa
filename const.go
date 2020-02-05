package ckassa

// RegMerchantPath путь для регистрации мерчанта.
const RegMerchantPath = "/registration/merchant"

// MobilePaymentPath путь для создание платежа в пользу мерчанта с баланса мобильного телефона.
const MobilePaymentPath = "/do/payment/mobile"

// PaymentPath путь для создания платежа в пользу мерчанта.
const PaymentPath = "/do/payment"

// CardTypes перечень доступных типов карт.
var CardTypes = [5]string{"undefined", "visa", "master_card", "maestro", "mir"}

// UserStates перечень доступных статусов пользователя.
var UserStates = [3]string{"active", "inactive", "disable"}

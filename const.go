package ckassa

// RegMerchantPath путь для регистрации мерчанта.
const RegMerchantPath = "/registration/merchant"

// MobilePaymentPath путь для создание платежа в пользу мерчанта с баланса мобильного телефона.
const MobilePaymentPath = "/do/payment/mobile"

// PaymentPath путь для создания платежа в пользу мерчанта.
const PaymentPath = "/do/payment"

// ReservePaymentPath путь для бронирования суммы на карте пользователя.
const ReservePaymentPath = "/do/payment-reserv"

// ConfirmReservePaymentPath путь для бронирования суммы на карте пользователя.
const ConfirmReservePaymentPath = "/provision-services/confirm"

// UpdatePayMerchantPath путь для обновления получателя ( мерчанта ) платежа.
const UpdatePayMerchantPath = "/update/payment/merchant"

// BalancePath путь для получения баланса мерчанта.
const BalancePath = "/get/merchant/wallet/balance"

// StatusPath путь для получения статуса мерчанта.
const StatusPath = "merchant/status"

// CardTypes перечень доступных типов карт.
var CardTypes = [5]string{"undefined", "visa", "master_card", "maestro", "mir"}

// UserStates перечень доступных статусов пользователя.
var UserStates = [3]string{"active", "inactive", "disable"}

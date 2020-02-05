package ckassa

// MERCHANT ==========================================

// MerchantRegPath путь для регистрации мерчанта.
const MerchantRegPath = "/registration/merchant"

// MerchantStatusPath путь для получения статуса мерчанта.
const MerchantStatusPath = "merchant/status"

// BalancePath путь для получения баланса мерчанта.
const BalancePath = "/get/merchant/wallet/balance"

// PAYMENT ===========================================

// MobilePaymentPath путь для создание платежа в пользу мерчанта с баланса мобильного телефона.
const MobilePaymentPath = "/do/payment/mobile"

// PaymentPath путь для создания платежа в пользу мерчанта.
const PaymentPath = "/do/payment"

// HOLD PAYMENT ======================================

// ReservePaymentPath путь для бронирования суммы на карте пользователя.
const ReservePaymentPath = "/do/payment-reserv"

// ConfirmReservePaymentPath путь для бронирования суммы на карте пользователя.
const ConfirmReservePaymentPath = "/provision-services/confirm"

// UpdatePayMerchantPath путь для обновления получателя ( мерчанта ) платежа.
const UpdatePayMerchantPath = "/update/payment/merchant"

// USER ==============================================

// UserRegPath путь для регистрации пользователя.
const UserRegPath = "/user/registration"

// UserStatusPath путь для получения статуса пользователя.
const UserStatusPath = "/user/status"

// CARD ==============================================

// CardRegPath путь для регистрации карты.
const CardRegPath = "/card/registration"

// CardListPath путь для получения списка карт.
const GetCardListPath = "/get/cards"

// CardDeactivationPath путь для деактивации зарегистрированной карты.
const CardDeactivationPath = " /card/deactivation"

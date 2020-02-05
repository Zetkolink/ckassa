package ckassa

const RegMerchantPath = "/registration/merchant"
const MobilePaymentPath = "/do/payment/mobile"

var CardTypes = [5]string{"undefined", "visa", "master_card", "maestro", "mir"}
var UserStates = [3]string{"active", "inactive", "disable"}

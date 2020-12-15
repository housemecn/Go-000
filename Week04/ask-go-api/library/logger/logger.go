package logger

// @Project: go-pay-api
// @Author: houseme
// @Description:
// @File: logger
// @Version: 1.0.0
// @Date: 2020/11/12 15:27
// @Package logger

const (
	OCR       = "ocr"
	Trade     = "trade"
	SMS       = "sms"
	Refund    = "refund"
	Notify    = "notify"
	Royalty   = "royalty"
	Auth      = "auth"
	Member    = "member"
	Grade     = "grade"
	Photocopy = "photocopy"
	Withdraw  = "withdraw"
	Article   = "article"
)

// GetOCRLogger Get OCR Logger
func GetOCRLogger() string {
	return OCR
}

// GetTradeLogger Get Trade Logger
func GetTradeLogger() string {
	return Trade
}

// GetSMSLogger Get SMS Logger
func GetSMSLogger() string {
	return SMS
}

// GetRefundLogger Get Refund Logger
func GetRefundLogger() string {
	return Refund
}

// GetNotifyLogger Get Notify Logger
func GetNotifyLogger() string {
	return Notify
}

// GetRoyaltyLogger Get Royalty Logger
func GetRoyaltyLogger() string {
	return Royalty
}

// GetAuthLogger Get Auth Logger
func GetAuthLogger() string {
	return Auth
}

// GetMemberLogger Get Member Logger
func GetMemberLogger() string {
	return Member
}

// GetGradeLogger Get Grade Logger
func GetGradeLogger() string {
	return Grade
}

// GetPhotocopyLogger Get Photocopy Logger
func GetPhotocopyLogger() string {
	return Photocopy
}

// GetWithdrawLogger Get Withdraw Logger
func GetWithdrawLogger() string {
	return Withdraw
}

// GetArticleLogger Get Article Logger
func GetArticleLogger() string {
	return Article
}

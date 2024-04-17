package defsms

import "time"

const (
	SmsLoginKeepTime       = time.Minute * 5
	SmsLoginLimitTime      = time.Second * 60
	SmsRenewPhoneKeepTime  = time.Minute * 5
	SmsRenewPhoneLimitTime = time.Second * 60
	SmsOldPhoneKeepTime    = time.Minute * 5
	SmsOldPhoneLimitTime   = time.Second * 60
)

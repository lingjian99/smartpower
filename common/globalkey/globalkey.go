package globalkey

var SmsCodeFormat = "sms:%s"

// 软删除
var DelStateNo int64 = 0  //未删除
var DelStateYes int64 = 1 //已删除

var DefalutTimeFormat = "2006-01-02 15:04:05"

var TimeFormatDay = "2006-01-02"
var TimeFormatMonth = "2006-01"

// 业务类型
const ServiceTypeAccount = 1 // 充值账户

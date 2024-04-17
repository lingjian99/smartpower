package xerr

const OK uint32 = 2000

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR uint32 = 100001

const SYS_RPC_COMMON_ERR uint32 = 200001

const (
	PARAM_ERROR      uint32 = 4000
	TOKEN_ERROR      uint32 = 100002
	DB_ERROR         uint32 = 100005
	REDIS_ERROR      uint32 = 100006
	SMS_SEND_ERROR   uint32 = 100007
	DATA_NOT_EXIST   uint32 = 100008
	DATA_EXIST       uint32 = 100009
	UserExists       uint32 = 101002
	UserNoPermission uint32 = 101003
	PhoneExists      uint32 = 101005
	EmailExists      uint32 = 101006
	RawPwdError      uint32 = 101007
	GroupExists      uint32 = 101008
	CollectorExists  uint32 = 101009
	InverterExists   uint32 = 101010
	ModelExists      uint32 = 101011
	ColSettingExists uint32 = 101012
	PowerNotExists   uint32 = 101013
	WrongTimeFormat  uint32 = 101014
	UserIsLocked     uint32 = 101015
	EnvNumIsExist    uint32 = 101016 //
	DeviceError      uint32 = 101017
	NoLockSuperUser  uint32 = 101018
	ApiLimit         uint32 = 101019
)

const (
	PayErrNo uint32 = 103001
)

var (
	errorCodeMap    = initErrorCode()
	UserNotExistErr = NewErrCode(101001)
	PhoneExistErr   = NewErrCode(101005)
	Passworderr     = NewErrCode(101004)
	SmsSendErr      = NewErrCode(SMS_SEND_ERROR)
	EmailExistErr   = NewErrCode(EmailExists)
	RawPwdErr       = NewErrCode(RawPwdError)
)

//var errorCodeMap map[uint32]string

func initErrorCode() map[uint32]string {
	// 初始化全局错误码
	errorCodeMap := make(map[uint32]string)
	errorCodeMap[2000] = "OK"
	/// 请求
	errorCodeMap[PARAM_ERROR] = "请求参数错误"

	errorCodeMap[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	errorCodeMap[DB_ERROR] = "数据库繁忙,请稍后再试"
	errorCodeMap[SMS_SEND_ERROR] = "短信发送失败"
	errorCodeMap[DATA_NOT_EXIST] = "数据不存在"
	errorCodeMap[DATA_EXIST] = "数据已存在"
	// Register
	errorCodeMap[101001] = "用户不存在"
	errorCodeMap[UserExists] = "用户已存在"
	errorCodeMap[UserNoPermission] = "用户没有权限"
	errorCodeMap[101004] = "密码错误"
	errorCodeMap[101005] = "手机号已经被注册"
	errorCodeMap[EmailExists] = "邮箱已被注册"
	errorCodeMap[RawPwdError] = "原密码错误"
	errorCodeMap[GroupExists] = "该工厂组名已存在"
	errorCodeMap[CollectorExists] = "数据采集器已在该工厂注册"
	errorCodeMap[InverterExists] = "逆变器已在该工厂注册"
	errorCodeMap[ModelExists] = "该型号的序号或名称已存在"
	errorCodeMap[ColSettingExists] = "该采集器已存在配置"
	errorCodeMap[PowerNotExists] = "该工厂不存在"
	errorCodeMap[UserIsLocked] = "账号已锁定"
	errorCodeMap[WrongTimeFormat] = "时间格式错误"
	errorCodeMap[NoLockSuperUser] = "无法锁定超级管理员"

	errorCodeMap[PayErrNo] = "支付失败"

	return errorCodeMap
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := errorCodeMap[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := errorCodeMap[errcode]; ok {
		return true
	} else {
		return false
	}
}

//func init() {
//	errorCodeMap = initErrorCode()
//}

package defcache

import "errors"

const (
	SmsLoginPrefix           = "km:sms:login:"             //登陆短信前缀
	SmsLoginLimitPrefix      = "km:sms:login_limit:"       //登陆短信限频前缀
	SmsRenewPhonePrefix      = "km:sms:renew_phone:"       //更换手机号短信前缀
	SmsRenewPhoneLimitPrefix = "km:sms:renew_phone_limit:" //更换手机号短信限频前缀
	SmsOldPhonePrefix        = "km:sms:old_phone:"         //验证旧手机号短信前缀
	SmsOldPHoneLimitPrefix   = "km:sms:old_phone_limit:"   //验证旧手机号短信限频前缀
	CliAvatarHistoryPrefix   = "km:cli:history_avatar:"    //客户端历史头像
	CliAvatarUserIdPrefix    = "km:cli:avatar:user_id:"    //用户头像对应的用户id
	OldPhoneVerifyCertify    = "km:old_phone:certify:"     //旧手机号验证通过标识
	ProvinceCityAreaCode     = "km:area_code"
	CliFriendCount           = "km:cli:friend_count:"  //用户好友数
	CliMchRecommend          = "km:cli:mch:recommend"  //商铺附近房间不为空的商铺
	WsRpcListenOn            = "km:e:ws_rpc_listen_on" //员工端rpc节点key
	ViolationImgPrefix       = "km:img:violation:"     //违规图片

	//*************Merchant*************//
	MchSmgLoginPrefix = "km:mch:sms:user:login:" //商户端用户登陆

	MchRoomImgHistoryPrefix    = "km:mch:history_room_img:"       //棋牌室房间历史图片
	MchRoomImgRoomIdPrefix     = "km:mch:room_img:room_id:"       // 棋牌室房间图片对应的房间id
	MchAvatarHistoryPrefix     = "km:mch:history_avatar:"         //商户端用户历史头像
	MchAvatarUserIdPrefix      = "km:mch:avatar:user_id:"         //商户头像对应的用户id
	MchProductImgHistoryPrefix = "km:mch:history_product_img:"    //商品历史图
	MchProductIdPrefix         = "km:mch:product_img:product_id:" //商品图对应的商品id
)

var (
	ErrCacheErrNotFound = errors.New("cache data not fount")
)

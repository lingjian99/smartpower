package tool

func PayMode(i int32) string {
	switch i {
	case 1:
		return "微信支付"
	case 2:
		return "支付宝支付"
	default:
		return ""
	}
}

func PayState(i int32) string {
	switch i {
	case -1:
		return "已取消"
	case 0:
		return "未支付"
	case 1:
		return "已支付"
	default:
		return ""
	}
}

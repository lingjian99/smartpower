package tool

import "fmt"

func MahjongType(typ int32) string {
	switch typ {
	case 0:
		return "常规麻将"
	case 1:
		return "杠头麻将"
	case 2:
		return "川式麻将"
	case 3:
		return "冲击麻将"
	default:
		return "未知类型"
	}
}

func MahjongDuration(d int32) string {
	if d >= 60 {
		return fmt.Sprintf("%d 小时 %d 分钟", d/60, d%60)
	}
	return fmt.Sprintf("%d 分钟", d)
}

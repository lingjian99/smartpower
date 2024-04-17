package uniqueid

import (
	"fmt"
	"smartpower/pkg/tool"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_THIRD_PAYMENT   SnPrefix = "MCT" //第三方支付流水记录前缀 /third_payment
	SN_PREFIX_BALANCE_PAYMENT SnPrefix = "ACT" // 余额支付
)

// 生成单号
func GenSn(snPrefix SnPrefix) string {
	return fmt.Sprintf("%s%s%s", snPrefix, time.Now().Format("20060102150405"), tool.Krand(8, tool.KC_RAND_KIND_NUM))
}

package protocol

// 硬件上传账单 len 32
type Billing struct {
	Keep0 byte
	//麻将类型
	//0 常规麻将
	//1 杠头麻将
	//2 川式麻将
	//3 冲击麻将
	MahKind   byte
	Keep1     byte
	StartTime [6]byte
	Duration  uint16
	//当前进行的击数/局数 XX击
	Count byte
	//老板每人每击点数抽成 XX点
	Rak byte

	Keep3 [13]byte
	Date  [6]byte
}

func NewBillingAck(success bool) []byte {
	var b byte
	if !success {
		b = 1
	}
	data := make([]byte, 16)
	data[0] = 0xca
	data[1] = 0x1b
	data[2] = 0x4d
	data[3] = 0x45
	data[4] = b

	return data
}

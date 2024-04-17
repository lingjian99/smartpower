package protocol

type UrlAck struct {
	MajorId byte
	MinorId byte
	BuildId uint32
	URL     [250]byte
}

func NewUrlAck(major, minor, buildId int64, url string) *UrlAck {

	u := &UrlAck{
		MajorId: byte(major),
		MinorId: byte(minor),
		BuildId: uint32(buildId),
		URL:     [250]byte{},
	}
	copy(u.URL[:], url)
	return u
}

// len = 16
type ServiceCmdReq struct {
	//0 呼叫
	ServiceType byte
	//呼叫发起者名称 GB2312 编码，每个汉字占两个字节 最多 12 字节
	Name [13]byte
	Keep [2]byte
}

//func (c *ServiceCmdReq) NameString() string {
//	s, _ := GBKToUTF8Bytes(c.Name[:])
//	return string(s)
//}

func (c *ServiceCmdReq) NameString() string {
	n2 := TrimNULL(c.Name[:])
	return GBKToUTF8String(n2)
}

var HeartbeatBody_16 = []byte{
	0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// HwWifi len = 64
type HwWifi struct {
	SSID    [32]byte
	Rssi    byte // 信号强度（RSSI）
	Quality byte // 信号质量
	IP      [4]byte
	Keep    [20]byte
	Date    [6]byte
}

type QueryArrears [16]byte

// 0x030b 功能码 len = 16
func NewQueryArrearsAck(count uint8) QueryArrears {
	//data := make([]byte, 16)
	var data QueryArrears
	data[0] = 0x2a
	data[1] = 0x4c
	data[2] = 0x5f
	data[3] = 0x13
	data[4] = count

	return data
}

func RawDataLen(length int, datas ...byte) []byte {
	ld := len(datas)
	if ld > length {
		ld = length
	}
	resp := make([]byte, length)
	for i := 0; i < ld; i++ {
		d := datas[i]
		resp[i] = d
	}
	return resp[:]
}

func RawData16(datas ...byte) []byte {
	l := len(datas)
	if l > 16 {
		l = 16
	}
	var resp [16]byte
	for i := 0; i < l; i++ {
		d := datas[i]
		resp[i] = d
	}
	return resp[:]
}

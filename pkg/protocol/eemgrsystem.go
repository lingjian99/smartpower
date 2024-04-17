package protocol

import "time"

// NewServerTimeSynchronization 同步时间 返回基于UTC偏移的时间
func NewServerTimeSynchronization(t time.Time, tz int) *Input {
	utc := t.UTC().Add(time.Hour * time.Duration(-tz))
	data := make([]byte, 0x10)
	data[0] = byte(utc.Year() - 2000)
	data[1] = byte(utc.Month())
	data[2] = byte(utc.Day())
	data[3] = byte(utc.Hour())
	data[4] = byte(utc.Minute())
	data[5] = byte(utc.Second())

	data[7] = 0x12
	data[8] = 0x2a
	data[9] = 0xb8
	data[10] = 0x7e
	return &Input{
		Head: NewHead(0x10, 0xFEFF, 0x02, 0x10),
		Data: data,
	}
}

// UAFFT 硬件发送FFT数据
type UAFFT struct {
	Data [32]uint16
	Keep [10]byte
	Date [6]byte
}

type IUPhaseAngle struct {
	YData [9]int16
	Keep  [8]byte
	Date  [6]byte
}

type RemoteControlRequest struct {
	OpCode uint16
	OpID   uint16
	Keep   [12]byte
}

type RemoteControlResult struct {
	//"操作反馈
	//0.成功
	//1.继电器操作失败
	//2.操作ID错误
	//3.工作模式错误
	//4.操作类型错误"
	Result byte
	Keep   [2]byte
	//"当前状态
	//0=运行状态
	//1=热备状态
	//2=冷备状态
	//3=检修状态
	RunningState byte
	//"主断路器状态
	//0=断开
	//1=闭合
	//3=无效"
	MainSwitch byte
	Keep1      [11]byte
}

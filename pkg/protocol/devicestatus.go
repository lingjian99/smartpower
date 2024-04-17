package protocol

// 硬件上传设备状态 len = 32
type HwStatus struct {
	//设备运行状态 0 空闲 1 正在进行
	RunningStatus byte
	//当前麻将类型
	//0 常规麻将
	//1 杠头麻将
	//2 川式麻将
	//3 冲击麻将
	MahKind byte
	//系统状态
	//0 正常
	//1 欠费锁定
	//2 系统故障
	SysLockStatus byte
	StartTime     [6]byte
	//麻将持续时间 XX分钟
	Duration uint16
	//当前进行的击数/局数 XX击
	Count byte
	//老板每人每击点数抽成 XX点
	Rak byte
	//环境温度 XX.X℃
	Temp         uint16
	DeviceVolume byte
	PlayersNum   byte
	Keep         [9]byte
	Date         [6]byte
}

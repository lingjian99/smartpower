package photovoltaic_pto

type DeviceInfo struct {
	Keep0      uint16   `json:"keep_0,omitempty"`
	DeviceId   uint16   `json:"device_id,omitempty"`
	DeviceType byte     `json:"device_type,omitempty"`
	Coding     byte     `json:"coding,omitempty"`
	BrandModel [42]byte `json:"brand_model,omitempty"`
}
type InverterData struct {
	Keep0                    uint16  `json:"keep_0,omitempty"`
	DeviceId                 uint16  `json:"device_id,omitempty"`
	PowerFactor              byte    `json:"power_factor,omitempty"`
	Status                   byte    `json:"status,omitempty"`
	Dc16                     uint16  `json:"dc_16,omitempty"`
	Dcv1                     uint16  `json:"dcv_1,omitempty"`
	Dc1                      uint16  `json:"dc_1,omitempty"`
	Dcv2                     uint16  `json:"dcv_2,omitempty"`
	Dc2                      uint16  `json:"dc_2,omitempty"`
	Dcv3                     uint16  `json:"dcv_3,omitempty"`
	Dc3                      uint16  `json:"dc_3,omitempty"`
	Dcv4                     uint16  `json:"dcv_4,omitempty"`
	Dc4                      uint16  `json:"dc_4,omitempty"`
	Dcv5                     uint16  `json:"dcv_5,omitempty"`
	Dc5                      uint16  `json:"dc_5,omitempty"`
	Dcv6                     uint16  `json:"dcv_6,omitempty"`
	Dc6                      uint16  `json:"dc_6,omitempty"`
	Dcv7                     uint16  `json:"dcv_7,omitempty"`
	Dc7                      uint16  `json:"dc_7,omitempty"`
	Dcv8                     uint16  `json:"dcv_8,omitempty"`
	Dc8                      uint16  `json:"dc_8,omitempty"`
	Dcv9                     uint16  `json:"dcv_9,omitempty"`
	Dc9                      uint16  `json:"dc_9,omitempty"`
	Dcv10                    uint16  `json:"dcv_10,omitempty"`
	Dc10                     uint16  `json:"dc_10,omitempty"`
	Ua                       uint16  `json:"ua,omitempty"`
	Ub                       uint16  `json:"uc,omitempty"`
	Uc                       uint16  `json:"ib,omitempty"`
	Ia                       uint16  `json:"ub,omitempty"`
	Ib                       uint16  `json:"ia,omitempty"`
	Ic                       uint16  `json:"ic,omitempty"`
	Uab                      uint16  `json:"uab,omitempty"`
	Ubc                      uint16  `json:"ubc,omitempty"`
	Uca                      uint16  `json:"ca,omitempty"`
	GridFrequency            uint16  `json:"grid_frequency,omitempty"`
	ActivePower              int32   `json:"active_power,omitempty"`
	ReactivePower            int32   `json:"reactive_power,omitempty"`
	ExistingPower            int32   `json:"existing_power,omitempty"`
	DailyPowerGeneration     uint32  `json:"daily_power_generation,omitempty"`
	MonthlyPowerGeneration   uint32  `json:"month_power_generation,omitempty"`
	AnnualGenerationCapacity uint32  `json:"annual-generation-capacity,omitempty"`
	TotalPowerGeneration     uint32  `json:"total_power_generation,omitempty"`
	Temperature              int16   `json:"temperature,omitempty"`
	PowerLimit               byte    `json:"power_limit,omitempty"`
	Keep2                    byte    `json:"keep_4,omitempty"`
	Dcv11                    uint16  `json:"dcv_11"`
	Dc11                     uint16  `json:"dc_11"`
	Dcv12                    uint16  `json:"dcv_12"`
	Dc12                     uint16  `json:"dc_12"`
	Dcv13                    uint16  `json:"dcv_13"`
	Dc13                     uint16  `json:"dc_13"`
	Dcv14                    uint16  `json:"dcv_14"`
	Dc14                     uint16  `json:"dc_14"`
	Dcv15                    uint16  `json:"dcv_15"`
	Dc15                     uint16  `json:"dc_15"`
	Dcv16                    uint16  `json:"dcv_16"`
	Date                     [6]byte `json:"date"`
}
type InverterFault struct {
	Keep0    uint16   `json:"keep_0,omitempty"`
	DeviceId uint16   `json:"device_id,omitempty"`
	Keep1    byte     `json:"keep_1,omitempty"`
	Coding   byte     `json:"coding,omitempty"`
	Content  [68]byte `json:"content,omitempty"`
	Date     [6]byte  `json:"date"`
}
type InvFaultV2 struct {
	Keep0                    uint16   `json:"keep_0,omitempty"`
	DeviceId                 uint16   `json:"device_id,omitempty"`
	InvStatus                byte     `json:"inv_status,omitempty"`
	Coding                   byte     `json:"coding,omitempty"`
	Content                  [68]byte `json:"content,omitempty"`
	Dc16                     uint16   `json:"dc_16,omitempty"`
	Dcv1                     uint16   `json:"dcv_1,omitempty"`
	Dc1                      uint16   `json:"dc_1,omitempty"`
	Dcv2                     uint16   `json:"dcv_2,omitempty"`
	Dc2                      uint16   `json:"dc_2,omitempty"`
	Dcv3                     uint16   `json:"dcv_3,omitempty"`
	Dc3                      uint16   `json:"dc_3,omitempty"`
	Dcv4                     uint16   `json:"dcv_4,omitempty"`
	Dc4                      uint16   `json:"dc_4,omitempty"`
	Dcv5                     uint16   `json:"dcv_5,omitempty"`
	Dc5                      uint16   `json:"dc_5,omitempty"`
	Dcv6                     uint16   `json:"dcv_6,omitempty"`
	Dc6                      uint16   `json:"dc_6,omitempty"`
	Dcv7                     uint16   `json:"dcv_7,omitempty"`
	Dc7                      uint16   `json:"dc_7,omitempty"`
	Dcv8                     uint16   `json:"dcv_8,omitempty"`
	Dc8                      uint16   `json:"dc_8,omitempty"`
	Dcv9                     uint16   `json:"dcv_9,omitempty"`
	Dc9                      uint16   `json:"dc_9,omitempty"`
	Dcv10                    uint16   `json:"dcv_10,omitempty"`
	Dc10                     uint16   `json:"dc_10,omitempty"`
	Ua                       uint16   `json:"ua,omitempty"`
	Ub                       uint16   `json:"uc,omitempty"`
	Uc                       uint16   `json:"ib,omitempty"`
	Ia                       uint16   `json:"ub,omitempty"`
	Ib                       uint16   `json:"ia,omitempty"`
	Ic                       uint16   `json:"ic,omitempty"`
	Uab                      uint16   `json:"uab,omitempty"`
	Ubc                      uint16   `json:"ubc,omitempty"`
	Uca                      uint16   `json:"ca,omitempty"`
	GridFrequency            uint16   `json:"grid_frequency,omitempty"`
	ActivePower              int32    `json:"active_power,omitempty"`
	ReactivePower            int32    `json:"reactive_power,omitempty"`
	ExistingPower            int32    `json:"existing_power,omitempty"`
	DailyPowerGeneration     uint32   `json:"daily_power_generation,omitempty"`
	MonthlyPowerGeneration   uint32   `json:"month_power_generation,omitempty"`
	AnnualGenerationCapacity uint32   `json:"annual-generation-capacity,omitempty"`
	TotalPowerGeneration     uint32   `json:"total_power_generation,omitempty"`
	Temperature              int16    `json:"temperature,omitempty"`
	PowerLimit               byte     `json:"power_limit,omitempty"`
	Keep2                    byte     `json:"keep_4,omitempty"`
	Dcv11                    uint16   `json:"dcv_11"`
	Dc11                     uint16   `json:"dc_11"`
	Dcv12                    uint16   `json:"dcv_12"`
	Dc12                     uint16   `json:"dc_12"`
	Dcv13                    uint16   `json:"dcv_13"`
	Dc13                     uint16   `json:"dc_13"`
	Dcv14                    uint16   `json:"dcv_14"`
	Dc14                     uint16   `json:"dc_14"`
	Dcv15                    uint16   `json:"dcv_15"`
	Dc15                     uint16   `json:"dc_15"`
	Dcv16                    uint16   `json:"dcv_16"`
	Keep3                    [12]byte `json:"keep_3"`
	Date                     [6]byte  `json:"date"`
}
type InvFaultV2Reply struct {
	Head1 byte    `json:"head_1"`
	Head2 byte    `json:"head_2"`
	Head3 byte    `json:"head_3"`
	Head4 byte    `json:"head_4"`
	Keep0 byte    `json:"keep_0"`
	Mid   uint16  `json:"mid"`
	Keep1 [9]byte `json:"keep_1"`
}
type PowerLimitSetting struct {
	Head1      byte
	Head2      byte
	Head3      byte
	ComNumber  byte
	Keep0      byte
	MidAddr    uint16
	Set        byte
	PowerLimit byte
	Keep1      [7]byte
}
type PowerLimitSettingReply struct {
	Head1       byte
	Head2       byte
	Head3       byte
	Head4       byte
	Status      byte
	MidAddr     uint16
	SettingType byte
	Keep0       [8]byte
}

type TestProto struct {
	Head1 byte
	Head2 byte
	Head3 byte
	Head4 byte
	ID    uint32
}

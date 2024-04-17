package photovoltaic_pto

type SettingReq struct {
	Head1 byte `json:"head_1"`
}
type SettingReply struct {
	Head1           byte     `json:"head_1"`
	Head2           byte     `json:"head_2"`
	Head3           byte     `json:"head_3"`
	ComNum          byte     `json:"com_num"`
	ModbusStartAddr byte     `json:"modbus_start_addr"`
	MidtAddr        uint16   `json:"midt_addr"`
	DeviceNum       byte     `json:"device_num"`
	Device1Model    byte     `json:"device_1_model"`
	Device2Model    byte     `json:"device_2_model"`
	Device3Model    byte     `json:"device_3_model"`
	Device4Model    byte     `json:"device_4_model"`
	Device5Model    byte     `json:"device_5_model"`
	Device6Model    byte     `json:"device_6_model"`
	Device7Model    byte     `json:"device_7_model"`
	Device8Model    byte     `json:"device_8_model"`
	Device9Model    byte     `json:"device_9_model"`
	Device10Model   byte     `json:"device_10_model"`
	Device11Model   byte     `json:"device_11_model"`
	Device12Model   byte     `json:"device_12_model"`
	Device13Model   byte     `json:"device_13_model"`
	Device14Model   byte     `json:"device_14_model"`
	Device15Model   byte     `json:"device_15_model"`
	Device16Model   byte     `json:"device_16_model"`
	Device17Model   byte     `json:"device_17_model"`
	Device18Model   byte     `json:"device_18_model"`
	Device19Model   byte     `json:"device_19_model"`
	Device20Model   byte     `json:"device_20_model"`
	Device21Model   byte     `json:"device_21_model"`
	Device22Model   byte     `json:"device_22_model"`
	Device23Model   byte     `json:"device_23_model"`
	Device24Model   byte     `json:"device_24_model"`
	Device25Model   byte     `json:"device_25_model"`
	Device26Model   byte     `json:"device_26_model"`
	Device27Model   byte     `json:"device_27_model"`
	Device28Model   byte     `json:"device_28_model"`
	Device29Model   byte     `json:"device_29_model"`
	Device30Model   byte     `json:"device_30_model"`
	Keep0           [26]byte `json:"keep_0"`
}

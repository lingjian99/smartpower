package photovoltaic_pto

type DataCollector struct {
	Year        byte    `json:"keep_0,omitempty"`
	Month       byte    `json:"keep_1,omitempty"`
	RSSI        byte    `json:"rssi"`
	Keep0       byte    `json:"keep0"`
	Temperature int16   `json:"temperature,omitempty"`
	Day         byte    `json:"day"`
	Hour        byte    `json:"hour"`
	Minute      byte    `json:"minute"`
	Second      byte    `json:"second"`
	Date        [6]byte `json:"date"`
}
type DataCollectInfo struct {
	Bid             [18]byte `json:"bid"`
	ProtocolVersion int16    `json:"protocol_version"`
	DeviceModel     [24]byte `json:"device_model"`
	DeviceUUID      int32    `json:"device_uuid"`
	SoftVersion1    byte     `json:"soft_version_1"`
	SoftVersion2    byte     `json:"soft_version_2"`
	SoftVersion3    int32    `json:"soft_version_3"`
	TimeZone        byte     `json:"time_zone"`
	Keep0           [5]byte  `json:"keep_0"`
}

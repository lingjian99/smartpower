package photovoltaic_pto

type EnvMonitorData struct {
	Keep0                   uint16  `json:"keep_0,omitempty"`
	DeviceId                uint16  `json:"device_id,omitempty"`
	ComponentTemperature    int16   `json:"component_temperature,omitempty"`      //组件温度
	EnvTemperature          int16   `json:"env_temperature,omitempty"`            //环境温度
	WindSpeed               uint16  `json:"wind_speed,omitempty"`                 //风速
	WindDirection           uint16  `json:"wind_direction,omitempty"`             //风向
	TotalRadiation          uint32  `json:"total_radiation,omitempty"`            //总辐射
	ScatteredRadiation      uint32  `json:"scattered_radiation,omitempty"`        //散辐射
	DirectRadiation         uint32  `json:"direct_radiation,omitempty"`           //直接辐射
	TotalRadiationDayDC     uint32  `json:"total_radiation_day_dc,omitempty"`     //总辐射日累积量
	ScatteredRadiationDayDC uint32  `json:"scattered_radiation_day_dc,omitempty"` //散辐射日累计量
	DirectRadiationDayDC    uint32  `json:"direct_radiation_day_dc,omitempty"`    //直接辐射日累积量
	Keep1                   [6]byte `json:"keep_1,omitempty"`
	Date                    [6]byte `json:"date"`
}

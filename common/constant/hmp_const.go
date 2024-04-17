package constant

const (
	TablePrimaryElectricity    = "hmp_primary_electricity"
	TableMainSwitchOperation   = "hmp_main_switch_operation"
	TableDehumi                = "hmp_dehumi"
	TableWaveRecordAlarm       = "hmp_wave_record_alarm_data"
	TableSystemData            = "hmp_system_sql"
	TableSecondaryElectricity  = "hmp_secondary_electricity"
	TableSwitchData            = "hmp_switch"
	TableMainSwitchPollingData = "hmp_main_switch_polling_data"
	TableSettingDataUpdateData = "hmp_setting_data_sql"
	TableDzHjsEit              = "hmp_dz_hjs_eit_sql"
)

///
///  11 除湿器数据
///  12 故障录波数据
///  13 定值数据
///  14 FFT 数据
var HtestCachedMap = map[string]HtestLayout{
	"gc_electrical":                          newLayout(3, "电气数据"),
	"lmp_electricity_polling_data":           newLayout(3, "低压进线管理机发送测量+保护数据"),
	"gc_electric_energy":                     newLayout(1, "电能数据"),
	"gc_date_section":                        newLayout(14, "计量时段数据"),
	"lmp_settingdata_update_data":            newLayout(3, "低压进线管理机发送保护定值数据"),
	"gc_fft_ua":                              newLayout(1, "A相电压FFT数据"),
	"gc_fft_ub":                              newLayout(1, "B相电压FFT数据"),
	"gc_fft_uc":                              newLayout(1, "C相电压FFT数据"),
	"gc_fft_ia":                              newLayout(1, "A相电流FFT数据"),
	"gc_fft_ib":                              newLayout(1, "B相电流FFT数据"),
	"gc_fft_ic":                              newLayout(1, "C相电流FFT数据"),
	"infft_polling_data":                     newLayout(1, "零序电流FFT数据"),
	"gc_switch":                              newLayout(1, "硬件发送开关量数据"),
	"lmp_system_polling_data":                newLayout(1, "低压进线管理机发送系统数据"),
	"gc_rvc_electrical_error":                newLayout(3, "电气报警数据"),
	"gc_rt_error":                            newLayout(1, "即时运行报警数据"),
	"hmp_primary_electricity_polling_data":   newLayout(3, "一次侧电气数据"),
	"hmp_secondary_electricity_polling_data": newLayout(3, "二次侧电气数据"),
	"hmp_switch_polling_data":                newLayout(2, "开入量数据"),
	"hmp_settingdata_update_data":            newLayout(13, "高压微机保护发送定值数据"),
	"hmp_dehumi_polling_data":                newLayout(11, "除湿器数据"),
	"hmp_waverecord_alarm_data":              newLayout(12, "故障录波数据"),
	"hmp_system_polling_data":                newLayout(1, "高压微机保护发送系统数据"),
	"hmp_mainswitch_operation_data":          newLayout(1, "断路器人工操作数据"),
	"hmp_main_switch_polling_data":           newLayout(1, "断路器操作统计数据"),
	"device_number":                          newLayout(1, "电气设备编号信息"),
	"gc_copper_temp":                         newLayout(2, "铜排温度数据"),
}

type HtestLayout struct {
	Layout uint16
	Name   string
}

func newLayout(layout uint16, name string) HtestLayout {
	return HtestLayout{
		Layout: layout,
		Name:   name,
	}
}

package constant

import (
	"strings"
)

const (
	// 1 "进线柜"
	CabModelIncoming = "1"
	// 2"进线总柜"
	CabModelIncomingMain = "2"
	//  3"计量柜"
	CabModelMetering = "3"
	// 4"压变柜"
	CabModelTrans = "4"
	// 5 "高压出线柜"
	CabModelHighOut = "5"
	// 6 "低压总柜"
	CabModelLow = "6"
	// 7"转折柜"
	CabModelTurn = "7"
	// 8"电容柜"
	CabModelCapacitor = "8"
	// 9 低压馈电柜
	CabModelLowFeeder = "9"
	// "10":"进线计量柜"
	CabModelMeasuring = "10"
	// "11":"进线隔离柜"
	CabModelIsolation = "11"
	// "12":"联络柜
	CabModelContact = "12"
	// "13": 低压计量柜
	CabModelLowMetering = "13"
)

var cabModelNameMap map[string]string

func init() {
	cabModelNameMap = make(map[string]string)
	cabModelNameMap[CabModelIncoming] = "高压进线柜"
	cabModelNameMap[CabModelIncomingMain] = "高压进线总柜"
	cabModelNameMap[CabModelMetering] = "计量柜"
	cabModelNameMap[CabModelTrans] = "高压压变柜"
	cabModelNameMap[CabModelHighOut] = "高压出线柜"
	cabModelNameMap[CabModelLow] = "低压进线柜"
	cabModelNameMap[CabModelTurn] = "转折柜"
	cabModelNameMap[CabModelCapacitor] = "低压电容柜"
	cabModelNameMap[CabModelLowFeeder] = "低压馈线柜"
	cabModelNameMap[CabModelMeasuring] = "低压进线计量柜"
	cabModelNameMap[CabModelIsolation] = "高压进线隔离柜"
	cabModelNameMap[CabModelContact] = "低压联络柜"
	cabModelNameMap[CabModelLowMetering] = "低压计量柜"
}

func HighTypeCab(mode string) bool {
	switch mode {
	case CabModelIsolation, CabModelIncoming, CabModelIncomingMain, CabModelTrans, CabModelHighOut, CabModelContact:
		return true
	}
	return false
}

func HighTypeBid(bid string) bool {
	return strings.HasPrefix(bid, "GYMP")
}

func CabModelString(mode string) string {
	return cabModelNameMap[mode]
}

const (
	TableformalFactoryBidsRelation = "formal_factory_bids_relation"
	TableSysAllAlert               = "sys_all_alert"
	TableFormalBillMaterials       = "formal_bill_materials"
	TableFormalSysStructure        = "formal_sys_structure"
	TableFormalFinalStructure      = "formal_final_sys_structure"
	TableFormalFacoties            = "formal_factories"
	TableCaOutage_period           = "ca_outage_period"
	TableGc_electricity            = "gc_electricity_ts"
	TableGc_electric_energy        = "gc_electric_energy"
	TableGc_electric_KxKnife       = "gc_electricity_kx"
	TableDevices                   = "devices"
	TableUser                      = "user"
	TableUsersFactories            = "users_factories"
	TableDyInfo                    = "dy_info"
	TableRVCElectricity            = "rvc_electricity"
	TableRVCCapacitance            = "rvc_capacitance"
	TableRVCSystem                 = "rvc_system"
	TableDeviceNumberData          = "device_number_data"
	TableGCRvcElectricalError      = "gc_rvc_electrical_error"
	TableMaxGcElectricalData1      = "gc_electricity_max_1"
	TableMaxGcElectDetailData      = "gc_electricity_max_detail"
	TableMaxGcElectricalData2      = "gc_electricity_max_2"
	TableGcElectricalStartTime     = "gc_electrical_starting"
	TableElectricityClassification = "electricity_classification"
	TableElectricityPrice2         = "electricity_price_2021"
	TableElectricityPrice          = "electricity_price"
	TableElectricityLossFactor     = "electricity_loss"
	TablePhaseAngleData            = "iu_phase_angle"
	TableGcSwitch                  = "gc_switch"
	TableDYKXSettingData           = "gc_dykx_setting_data"
	TableGenAlarmData              = "gen_alarm_data"
	TableStructHistory             = "struct_history"
	TableFrozenEventData           = "frozen_event_data"
	TableGYYBElePollingData        = "gyyb_ele_polling_data"
	TableGYSystemPollingData       = "gyyb_system_polling_data"
	TableGYYBSettingData           = "gyyb_settingdata_polling_data"
	TableGYGLSwitchData            = "gygl_switch_polling_data"
	//
	TableRVCelectricityPollingData = "rvc_electricity_data"
	TableRVCswitch_polling_data    = "rvc_switch_data"
	TableRVCsystem_polling_data    = "rvc_system_data"
	TableRVCsetting_data           = "rvc_setting_data"
	TableRVCcapcity_polling_data   = "rvc_capcity_data"
	TableRVCcos_polling_data       = "rvc_cos_data"
	TableRVCcos_polling_data2      = "rvc_cos_data_year" // 新增 year_id

	//
	TableLMPOLPStartData            = "lmp_olp_start"
	TableOperationDataSummary       = "operation_summary_data"
	TableOperationDataStep          = "operation_step_data"
	TableOperationDataEnd           = "operation_end_data"
	TableAutoOverloadProtectionData = "lmp_auto_protection_data"

	// question
	TableQuestionTagRelation  = "question_tag_relation"
	TableQuestionHotspot      = "question_hotspot"
	TableQuestionRoutine      = "question_routine"
	TableQuestionTag          = "question_tag"
	TableQuestionHotspotReply = "question_hotspot_reply"
	// dev_ops
	TableDevOps                     = "devops"
	TableInfoOutageNoc              = "info_outage"              // 停电通知表
	TableInfoOutageRead             = "info_outage_read"         // 停电通知已读
	TableInfoPowerSupplyReliability = "info_service_reliability" // 年供电可靠性

	// can 线 网络部署
	TableCommCan              = "comm_can"
	TableCommTransmissionLine = "comm_transmission_line"

	/// 处理过的数据
	TableProcessedEnergyDay = "pr_gc_energy_day"     // TODO 删除
	TableEnergyDailyCost    = "pr_energy_daily_cost" // TODO 删除
	// TableProcessedEnergyMonth   = "pr_gc_energy_month"
	TableProcessedCollectEnergy      = "pr_collect_energy"
	TableProcessedCollectEnergyDaily = "pr_collect_energy_daily"
	// TableEnergyFor2021          = "pr_energy_start_2021"
	//
	TableToolPowerFactor = "tool_power_factor"
)

const (
	// "2006-01-02 15:04:05"
	TimeFormat string = "2006-01-02 15:04:05"
	// TimeDayFormat "2006-01-02"
	TimeDayFormat string = "2006-01-02"
	// TimeFormatTZ DB "2006-01-02T15:04:05Z"
	TimeFormatTZ string = "2006-01-02T15:04:05Z"
	// TimeFormatAlarmDate 告警
	TimeFormatAlarmDate string = "2006.01.02 15:04:05"

	TimeMonthFormat string = "2006-01"
	TimeYearFormat  string = "2006"
)

const (
	EleDefaultGreen = 0 // 默认，有涌动效果
	EleYellow       = 1 // 告警 80-100
	EleRed          = 2 // 过载 > 100
	EleWhite        = 3 // 无电流无电压 白色
	EleStaticGreen  = 4 // 有电压静止无涌动
	EleStaticRed    = 6
)

const (
	IaType  = 1
	IbType  = 2
	IcType  = 3
	InType  = 4
	KIaType = 11
	KIbType = 12
	KIcType = 13
)

const (
	HighHighAccountForm = 0 // 高供高计
	HighLowAccountForm  = 1 // 高供低计
	SpotPowerPrice      = 0 // 电度电价
	TimeOfUsePrice      = 1 // 分时电价
)

const (
	LinkTypeMaster = "master"
	LinkTypeBranch = "branch"
)

const (
	RedisKeyStarting = "start_time"
)

// 特殊管理机。使用 GYYB 相关协议
const MultifunctionBid = "GYGL2011100930500X"

const (
	PrefixGYGL string = "GYGL"
	PrefixGYYB string = "GYYB"
	PrefixGYMP string = "GYMP"
	PrefixDYMP string = "DYMP"
	PrefixDYKX string = "DYKX"
	PrefixDRVC string = "DRVC"
	PrefuxRVC  string = "RVC"
)

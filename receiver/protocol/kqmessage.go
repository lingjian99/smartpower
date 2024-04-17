package protocol

import (
	"fmt"
	"strings"
)

type KqMessage struct {
	Bid     string
	Code    uint16 //功能码
	Flag    byte   //
	DataDst byte   // 数据方向
	Data    []byte
}

func (msg *KqMessage) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("[%#v %x %x][%s]", msg.Code, msg.Flag, msg.DataDst, msg.Bid))
	if len(msg.Data) == 0 {
		sb.WriteString("[]")
	} else {
		sb.WriteByte('[')
		for _, d := range msg.Data {
			sb.WriteString(fmt.Sprintf("%x ", d))
		}
		sb.WriteByte(']')
	}
	return sb.String()
}

const (
	CodeRegister0x01ff           = 0x01ff // 硬件上线注册
	CodeWifi0xf9ff               = 0xf9ff
	CodeState0x010b              = 0x010b
	CodeOrder0x020b              = 0x020b
	CodeAskArrears0x030b         = 0x030b
	CodeAskUpdate0x040b          = 0x040b
	CodeService0x050b            = 0x050b
	CodeDeviceInfo0x030c         = 0x030c //上传硬件信息
	CodeInverterData0x010c       = 0x010c //逆变器实时数据
	CodeInverterFault0x020c      = 0x020c //逆变器异常信息
	CodeEnvMonitor0x040c         = 0x040c // 环境监测仪实时数据
	CodeDataCollectStatus0x040c  = 0x050c //数据采集器状态
	CodeServerTelemetryReq0x060c = 0x060c //服务器遥测请求
	CodeDeviceSetting0x070c      = 0x070c //设备配置请求
	CodePowerLimit0x080c         = 0x080c // 功率限制
	CodeInvFault0x090c           = 0x090c //异常信息新协议
	CodeTest0xa0c                = 0xa0c  //测试协议
	//case 0xf9ff: // 上传 Wifi 信息
	//return l.hardwarewifi0xF9ff(in)
	//case 0x010b: // 上传硬件设备状态
	//return l.hardwareStatus0x010b(in)
	//case 0x020b: // 上传账单
	//return l.hardwareBilling0x020b(in)
	//case 0x030b:
	//return l.hardwareArrearsStatus0x030b(in)
	//case 0x040b:
	//return l.hardwareUpdating0x040b(in)
	//case 0x050b:
	//return l.hardware0x050b(in)
)

// 工业数据协议-通用协议
const (
	CodeUAFFT0x05ff                     = 0x05ff // 硬件发送FFT数据
	CodeServerMaterialListRequest0x0EFF = 0x0EFF // 服务器请求材料清单数据

	CodeHardwareelectricityalarmruntime0x12FF = 0x12FF //硬件发送即时电气报警数据
	CodeHardwaresystemalarmruntime0x13FF      = 0x13FF //硬件发送即时运行报警数据
	CodeHardwareELECLASS0x14ff                = 0x14ff //硬件上传电气层级 -- device_number
	CodeHMPDehumi0x15ff                       = 0x15ff //发送除湿器数据
	CodeIUphaseangle0x17ff                    = 0x17ff //硬件上传相位数据
	CodeHardwareOperationDataSummary0x18ff    = 0x18ff // 操作指南数据
	CodeHardwareFrozenData0x20ff              = 0x20ff // 冻结数据
	CodeGCEMpollingdata0x02FF                 = 0x02FF // 管理机发送电能数据
	CodeGCEMTimeSlot0x03FF                    = 0x03FF // 管理机发送计量时段数据

	CodeGYYB_ELE_polling_data0x0205    = 0x0205 // 发送模拟量+开入量
	CodeGYYB_System_polling_data0x1005 = 0x1005 // 发送系统数据
	//CodeGYYB_System_polling_data0x1006      = 0x1006 // 发送系统数据
	CodeGYYB_Settingdata_polling_data0x0405 = 0x0405 // 发送定值数据
	CodeGYGL_SWITCH_polling_data0x0206      = 0x0206 // 发送开入量

	CodeHMP_primary_electricity0x0202       = 0x0202 // 高压微机保护发送一次侧电气数据
	CodeHMP_secondary_electricity0x0302     = 0x0302
	CodeHMP_Switch_polling_data0x0b02       = 0x0b02
	CodeHMP_Settingdata_Update_data0x0402   = 0x0402
	CodeHMP_MainSwitch_polling_data0x1102   = 0x1102
	CodeHMP_MainSwitch_operation_data0x1702 = 0x1702
	HMP_WaveRecord_alarm_data0x1602         = 0x1602
	CodeHMP_system_polling_data0x1002       = 0x1002

	//高压小高配进出线
	CodeGXGP_Settingdata_polling0x0407 = 0x0407
	CodeGXGP_IO_polling0x0b07          = 0x0b07

	//
	CodeLMP_electricity_polling_data0x0203 = 0x0203 // 低压进线管理机发送测量+保护数据
	CodeLMP_Settingdata_Update_data0x0403  = 0x0403
	CodeLMP_system_polling_data0x1003      = 0x1003
	LMP_OLP_START0x1803                    = 0x1803

	CodeGC_electricity_polling_data0x0201 = 0x0201
	CodeGC_Switch_polling_data0x0b01      = 0x0b01
	CodeGC_CopperTemp_polling_data0x0d01  = 0x0d01
	CodeDYJXGC_system_polling_data0x1001  = 0x1001
	CodeDYJXGC_MainSwitch_polling0x1101   = 0x1101
	CodeDYJXGC_MainSwitch_operation0x1701 = 0x1701
	CodeDYJXGC_Setting_0x1801             = 0x1801

	//
	CodeRVC_Electricity_polling_data0x0204 = 0x0204
	CodeRVC_Settingdata_polling_data0x0404 = 0x0404
	CodeRVC_Switch_polling_data0x0304      = 0x0304
	CodeRVC_System_polling_data0x1004      = 0x1004
	CodeRVC_Capcity_polling_data0x0504     = 0x0504
	CodeRVC_Cos_polling_data0x0604         = 0x0604
)

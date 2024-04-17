package photovoltaic_pto

type PullDeviceData struct {
	Head1    byte
	Head2    byte
	Head3    byte
	Head4    byte
	DeviceId uint16
	Keep0    [10]byte
}

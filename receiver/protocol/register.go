package protocol

// BodyLen = 64

type Register struct {
	Bid     [18]byte
	Version uint16
	Model   [24]byte
	//SysInfo [31]byte
	UUID    [8]byte // mac
	MajorId byte
	MinorId byte
	BuildId uint32
	//Date    [6]byte
	Timezone int8 // -11~+12
	Keep     [5]byte
}

type RegisterResp struct {
	Data [58]byte
	Date [6]byte
}

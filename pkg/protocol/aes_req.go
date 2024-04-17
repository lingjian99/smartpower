package protocol

// AESReq BodyLen = 0x0080
type AESReq struct {
	AES [16]byte

	//AES加密类型
	//0 ECB
	//1 CBC
	//2 CFB
	//3 CTR
	//4 OFB
	CType uint8
	//"AES加密长度
	//0 128bits
	//1 192bits
	//2 256bits"
	CLen uint8
}

//func (r *AESReq) Unpack(data []byte) error {
//	copy(r.AES[:], data[5:15])
//	r.CType = data[16]
//	r.CLen = data[17]
//	return nil
//}

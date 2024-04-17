package defimage

import "time"

const (
	CliAvatarDir  = "client/avatar"
	CliAvatarType = "avatar"
	MchRoomImgDir = "mch/room"
	MchAvatarDir  = "mch/avatar"
	MchProductDir = "mch/product"
)

const (
	CliAvatarExpire = time.Minute * 1
)

var sourceMap = map[string]string{
	CliAvatarDir:  "用户头像",
	MchRoomImgDir: "房间图片",
	MchAvatarDir:  "用户头像",
}

func ImgSource(dir string) string {
	s, ok := sourceMap[dir]
	if !ok {
		return "其他"
	}
	return s
}

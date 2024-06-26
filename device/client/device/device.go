// Code generated by goctl. DO NOT EDIT.
// Source: device.proto

package device

import (
	"context"

	"smartpower/device/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BidReq              = pb.BidReq
	Conn                = pb.Conn
	DeviceData          = pb.DeviceData
	DeviceOnline        = pb.DeviceOnline
	EMTimeSlot          = pb.EMTimeSlot
	EMTimeSlot_TimeSlot = pb.EMTimeSlot_TimeSlot
	Empty               = pb.Empty
	RemoteControlReq    = pb.RemoteControlReq
	RemoteControlResp   = pb.RemoteControlResp

	Device interface {
		AnaDeviceConnected(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DeviceOnline, error)
		SendToDevice(ctx context.Context, in *DeviceData, opts ...grpc.CallOption) (*Empty, error)
		CloseConn(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Empty, error)
		// 服务器请求材料清单数据
		Server_MaterialList_Request(ctx context.Context, in *BidReq, opts ...grpc.CallOption) (*Empty, error)
		// 请求远程控制
		RemoteControlRequest(ctx context.Context, in *RemoteControlReq, opts ...grpc.CallOption) (*RemoteControlResp, error)
		// 服务器下发时段配置数据 载体固定96字节
		EMTimeSlotPollingPush(ctx context.Context, in *EMTimeSlot, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultDevice struct {
		cli zrpc.Client
	}
)

func NewDevice(cli zrpc.Client) Device {
	return &defaultDevice{
		cli: cli,
	}
}

func (m *defaultDevice) AnaDeviceConnected(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DeviceOnline, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.AnaDeviceConnected(ctx, in, opts...)
}

func (m *defaultDevice) SendToDevice(ctx context.Context, in *DeviceData, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.SendToDevice(ctx, in, opts...)
}

func (m *defaultDevice) CloseConn(ctx context.Context, in *Conn, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.CloseConn(ctx, in, opts...)
}

// 服务器请求材料清单数据
func (m *defaultDevice) Server_MaterialList_Request(ctx context.Context, in *BidReq, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.Server_MaterialList_Request(ctx, in, opts...)
}

// 请求远程控制
func (m *defaultDevice) RemoteControlRequest(ctx context.Context, in *RemoteControlReq, opts ...grpc.CallOption) (*RemoteControlResp, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.RemoteControlRequest(ctx, in, opts...)
}

// 服务器下发时段配置数据 载体固定96字节
func (m *defaultDevice) EMTimeSlotPollingPush(ctx context.Context, in *EMTimeSlot, opts ...grpc.CallOption) (*Empty, error) {
	client := pb.NewDeviceClient(m.cli.Conn())
	return client.EMTimeSlotPollingPush(ctx, in, opts...)
}

package logic

import (
	"context"
	"smartpower/receiver/internal/svc"
)

func DeviceConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) func(string, []byte) error {

	return func(bid string, in []byte) error {
/* 		body := new(protocol.QueryArrears)
		err := encodingx.Unmarshal(in, body)
		if err != nil {
			logx.Errorf("encodingx.Unmarshal: err%+v", err)
			return nil
		}

		c, err := svcCtx.McOrderModel.CountOrderByBid(ctx, bid, 0)
		if err != nil {
			logx.Errorf("CountOrderByBid: err: %+v", err)
		}
		logx.Debugf("BID:%s ordercount: %d", bid, c)
		data := protocol.NewQueryArrearsAck(uint8(c))

		_, err = svcCtx.HwRpc.SendToDevice(ctx, &device.DeviceData{
			Bid:  bid,
			Code: 0x030b,
			Flag: 0x02,
			Dst:  0x10,
			Data: data[:],
		})

		if err != nil {
			logx.Errorf("AskArrearsHandle %v", err)
		} */

		return nil
	}
}

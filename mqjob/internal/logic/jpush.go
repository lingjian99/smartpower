package logic

import (
	"context"
	"smartpower/common/jpush"
	"smartpower/mqjob/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

func PushMsg(ctx context.Context, svcCtx *svc.ServiceContext, alert string, rids []string, extras map[string]interface{}) error {
	appkey := svcCtx.Config.JiPushConf.Appkey
	masterSecret := svcCtx.Config.JiPushConf.MasterSecret

	var pf jpush.Platform
	pf.All()
	// Audience:
	var at jpush.Audience
	at.SetID(rids)

	//extras := map[string]interface{}{
	//	"type":      jobtype.NotifycationTypeServiceCall,
	//	"alert":     alert,
	//	"eventTime": time.Now().Format(globalkey.DefalutTimeFormat),
	//}
	// Notification
	var n jpush.Notification
	n.SetAlert(alert)
	n.SetAndroid(&jpush.AndroidNotification{
		AlertType:         7,
		DisplayForeground: "1",
		Alert:             alert, Title: "", Extras: extras})
	n.SetIos(&jpush.IosNotification{Alert: alert, Extras: extras})

	payload := jpush.NewPayLoad()
	payload.SetPlatform(&pf)
	payload.SetAudience(&at)
	payload.SetNotification(&n)
	//payload.SetMessage(&m)
	c := jpush.NewJPushClient(appkey, masterSecret) // appKey and masterSecret can be gotten from https://www.jiguang.cn/
	data, err := payload.Bytes()
	if err != nil {
		logx.WithContext(ctx).Errorf("ProcessTask MsgUnpaidNotifyUser, err :%+v", err)
		return err
	}
	res, err := c.Push(data)
	if err != nil {
		logx.Info(err)
	} else {
		logx.Info(res)
	}
	return err
}

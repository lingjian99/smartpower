package aliyunsms

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/pkg/errors"
)

type SmsConfig struct {
	AccessKey string
	Secret    string
}

func SendSms(tel, code string, conf SmsConfig) error {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", conf.AccessKey, conf.Secret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = tel             //接收短信的手机号码
	request.SignName = "宁波大家小家网络科技"        //短信签名名称
	request.TemplateCode = "SMS_193115749" //短信模板ID
	request.TemplateParam = "{\"code\":" + code + "}"

	response, err := client.SendSms(request)
	if err != nil {
		return err
	}
	if response.Code == "OK" {
		return nil
	}
	return errors.New(response.Message)
}

//func SendAliyunSms2(tel string, code string) error {
//client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", beego.AppConfig.String("aliyunaccesskey"), beego.AppConfig.String("accessKeySecret"))
//request := requests.NewCommonRequest()
//request.Method = "POST"
//request.Scheme = "https" // https | http
//request.Domain = "dysmsapi.aliyuncs.com"
//request.Version = "2017-05-25"
//request.ApiName = "SendSms"
//request.QueryParams["RegionId"] = "cn-hangzhou"
//request.QueryParams["PhoneNumbers"] = tel                        //手机号
//request.QueryParams["SignName"] = "宁波大家小家网络科技"                   //阿里云验证过的项目名 自己设置
//request.QueryParams["TemplateCode"] = "SMS_193115749"            //阿里云的短信模板号 自己设置
//request.QueryParams["TemplateParam"] = "{\"code\":" + code + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
//response, err := client.ProcessCommonRequest(request)
////fmt.Print(client.DoAction(request, response))
////  fmt.Print(response)
//if err != nil {
//	return errors.Wrapf(xerr.NewErrCode(xerr.SMS_SEND_ERROR), "err: %v", err)
//}
//type Message struct {
//	Message   string
//	RequestId string
//	BizId     string
//	Code      string
//}
//var message Message //阿里云返回的json信息对应的类
////记得判断错误信息
//re := json.Unmarshal(response.GetHttpContentBytes(), &message)
//beego.Info(re)
//if message.Message != "OK" {
//	log.Info("发送短信失败" + message.Message)
//	return errors.Wrapf(xerr.NewErrCode(xerr.SMS_SEND_ERROR), "msg: %v", message)
//}
//return nil
//}

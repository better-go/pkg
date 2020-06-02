package jpush

/*
ref: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/

	- https://github.com/ylywyn/jpush-api-go-client
		- 项目活跃
	- https://github.com/xialeistudio/go-jpush
	- https://github.com/jpush/jpush-api-python-client

	- usage:
		- https://github.com/jpush/jmessage-api-python-client/blob/master/example/messages/send_message.py

	- android test:
		- https://docs.jiguang.cn/jpush/client/Android/android_3m/
		- 客户端获取: RegistrationID
			- https://docs.jiguang.cn/jpush/client/Android/android_api/#registrationid-api
*/

import (
	"pkg/log"

	pushSDK "github.com/ylywyn/jpush-api-go-client"
)

var (
	// 支持的推送平台
	AllPlatform = "" // 所有平台
	IOS         = pushSDK.IOS
	Android     = pushSDK.ANDROID
	WinPhone    = pushSDK.WINPHONE
)

type JPush struct {
	// conf:
	AppKey    string // JPush key
	AppSecret string

	// push client:
	client *pushSDK.PushClient
}

func NewJPush(appKey string, appSecret string) *JPush {
	return &JPush{
		AppKey:    appKey,
		AppSecret: appSecret,
		client:    pushSDK.NewPushClient(appSecret, appKey),
	}
}

// 设置被推送的设备平台: iOS/Android 等
func (m *JPush) SetPlatform(platform string) *pushSDK.Platform {
	pf := new(pushSDK.Platform)

	// check:
	if platform == AllPlatform {
		pf.All()
	} else {
		pf.Add(platform)
	}

	return pf
}

// 设置被推送范围: 用户ID/用户 tag 等
func (m *JPush) SetAudience(isAll bool, ids []string, tags []string, tagAnds []string, alias []string) *pushSDK.Audience {
	ad := new(pushSDK.Audience)

	// all
	if isAll {
		ad.All()
		return ad
	}

	// registration_id
	if len(ids) != 0 {
		ad.SetID(ids)
	}

	// tag
	if len(tags) != 0 {
		ad.SetTag(tags)
	}

	// tag_and
	if len(tagAnds) != 0 {
		ad.SetTagAnd(tagAnds)
	}

	// alias
	if len(alias) != 0 {
		ad.SetAlias(alias)
	}

	return ad
}

// 设置推送弹出提示:
func (m *JPush) SetNotice(alertText string, androidText string, iOSText string, winPhoneText string) *pushSDK.Notice {
	n := new(pushSDK.Notice)
	n.SetAlert(alertText)

	if androidText != "" {
		n.SetAndroidNotice(&pushSDK.AndroidNotice{
			Alert:     androidText,
			Title:     "",
			BuilderId: 0,
			Extras:    nil,
		})
	}

	// ios:
	if iOSText != "" {
		n.SetIOSNotice(&pushSDK.IOSNotice{
			Alert:            iOSText,
			Sound:            "",
			Badge:            "",
			ContentAvailable: false,
			MutableContent:   false,
			Category:         "",
			Extras:           nil,
		})
	}

	//
	if winPhoneText != "" {
		n.SetWinPhoneNotice(&pushSDK.WinPhoneNotice{
			Alert:    winPhoneText,
			Title:    "",
			OpenPage: "",
			Extras:   nil,
		})
	}

	return n
}

// 设置推送消息报文内容:
func (m *JPush) SetMessage(title string, content string) *pushSDK.Message {
	msg := new(pushSDK.Message)
	msg.Title = title
	msg.Content = content
	return msg
}

// 推送完整报文初始化:
func (m *JPush) SetPayload(platform *pushSDK.Platform, audience *pushSDK.Audience, notice *pushSDK.Notice, message *pushSDK.Message) *pushSDK.PayLoad {
	p := pushSDK.NewPushPayLoad()

	p.SetPlatform(platform)
	p.SetAudience(audience)
	p.SetNotice(notice)
	p.SetMessage(message)
	return p
}

// 发送消息:
func (m *JPush) SendMessage(payload *pushSDK.PayLoad) error {
	bytes, err := payload.ToBytes()
	log.Debugf("jpush payload: %+v", string(bytes))
	if err != nil {
		log.Errorf("jpush payload format error, err=%v", err)
		return err
	}

	// do push:
	result, err := m.client.Send(bytes)
	log.Debugf("jpush send done, result=%+v, err=%v", result, err)
	return err
}

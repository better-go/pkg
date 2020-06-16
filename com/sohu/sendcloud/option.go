package sendcloud

const (
	// send cloud api host:
	urlMail         = "http://api.sendcloud.net/apiv2/mail/send"         // 邮件普通发送
	urlMailTemplate = "http://api.sendcloud.net/apiv2/mail/sendtemplate" // 邮件模板发送
	urlSms          = "http://www.sendcloud.net/smsapi/send"             // 短信普通发送
	urlSmsVoice     = "http://www.sendcloud.net/smsapi/sendVoice"        // 短信语音发送
)

//
type Option struct {
	ApiUser string // email: user
	ApiKey  string // email: key
	SmsUser string // sms: user
	SmsKey  string // sms: key
}

package sendcloud

import (
	"net/url"
	"strings"
)

/*
api:
	- https://www.sendcloud.net/doc/email_v2/
		- 发送: https://www.sendcloud.net/doc/email_v2/send_email/
	- https://www.sendcloud.net/doc/sms/

*/

const (
	// send cloud api host:
	urlMail         = "http://api.sendcloud.net/apiv2/mail/send"         // 邮件普通发送
	urlMailTemplate = "http://api.sendcloud.net/apiv2/mail/sendtemplate" // 邮件模板发送
	urlSms          = "http://www.sendcloud.net/smsapi/send"             // 短信普通发送
	urlSmsVoice     = "http://www.sendcloud.net/smsapi/sendVoice"        // 短信语音发送
)

//
type Option struct {
	EmailApiUser string // email: user
	EmailApiKey  string // email: key
	SmsApiUser   string // sms: user
	SmsApiKey    string // sms: key
}

///////////////////////////////////////////////////////////////////////////////////////////////////

// Mail email content: https://www.sendcloud.net/doc/email_v2/send_email/
type Mail struct {
	// key:
	ApiUser string `json:"apiUser"` // API_USER
	ApiKey  string `json:"apiKey"`  // API_KEY

	// address:
	AddrFrom string   `json:"from"` // 发件人地址
	AddrTo   []string `json:"to"`   // 收件人地址. 多个地址使用';'分隔, 如 ben@ifaxin.com;joe@ifaxin.com
	AddrCc   []string `json:"cc"`   // 抄送地址. 多个地址使用';'分隔
	AddrBcc  []string `json:"bcc"`  // 密送地址. 多个地址使用';'分隔

	//
	FromName string   `json:"fromName"` // 发件人名称. 显示如: ifaxin客服支持<support@ifaxin.com>
	ReplyTo  []string `json:"replyTo"`  // 设置用户默认的回复邮件地址.多个地址使用';'分隔，地址个数不能超过3个. 如果 replyTo 没有或者为空, 则默认的回复邮件地址为 from

	// content:
	Subject     string   `json:"subject"`     // 标题
	Html        string   `json:"html"`        // 邮件的内容. 邮件格式为 text/html
	Plain       string   `json:"plain"`       // 邮件的内容. 邮件格式为 text/plain
	Headers     string   `json:"headers"`     // 邮件头部信息. JSON 格式, 比如:{"header1": "value1", "header2": "value2"}
	Attachments []string `json:"attachments"` // 邮件附件. 发送附件时, 必须使用 multipart/form-data 进行 post 提交 (表单提交)

	// switch:
	UseNotification bool `json:"useNotification"` // 默认值: false. 是否使用回执
	UseAddressList  bool `json:"useAddressList"`  // 默认值: false. 是否使用地址列表发送
}

func (m *Mail) Default(apiUser string, apiKey string) {
	m.ApiUser = apiUser
	m.ApiKey = apiKey
}

func (m *Mail) ToSend() url.Values {
	v := url.Values{}

	v.Add("apiUser", m.ApiUser)
	v.Add("apiKey", m.ApiKey)
	v.Add("from", m.AddrFrom)
	v.Add("to", strings.Join(m.AddrTo, ";"))
	v.Add("subject", m.Subject)
	if m.Html != "" {
		v.Add("html", m.Html)
	}
	if m.Plain != "" {
		v.Add("plain", m.Plain)
	}

	return v
}

///////////////////////////////////////////////////////////////////////////////////////////////////

// 模板邮件
type MailTemplate struct {
	// key:
	ApiUser string `json:"apiUser"` // API_USER
	ApiKey  string `json:"apiKey"`  // API_KEY

	//
}

func (m *MailTemplate) Default(apiUser string, apiKey string) {
	m.ApiUser = apiUser
	m.ApiKey = apiKey
}

func (m *MailTemplate) ToSend() url.Values {
	v := url.Values{}

	v.Add("apiUser", m.ApiUser)
	v.Add("apiKey", m.ApiKey)

	return v
}

///////////////////////////////////////////////////////////////////////////////////////////////////

// Message sms message
type Message struct {
}

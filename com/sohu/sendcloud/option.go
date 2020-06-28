package sendcloud

import (
	"net/url"
	"strings"

	"github.com/better-go/pkg/convert"
	"github.com/better-go/pkg/time"
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

	// message type: 0表示短信, 1表示彩信,2表示国际短信,3表示国内语音,5表示影音 默认值为0
	MessageTypeCN      = 0 // 短信, 默认值为0
	MessageTypeOversea = 2 // 国际短信
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

// Message sms message: https://www.sendcloud.net/doc/sms/api/#send
type Message struct {
	SmsUser string `json:"smsUser"` //

	//
	TemplateID int64    `json:"templateId"` // 短信模板ID
	LabelID    int64    `json:"labelId"`    // 短信标签ID
	MsgType    int64    `json:"msgType"`    // 0表示短信, 1表示彩信,2表示国际短信,3表示国内语音,5表示影音 默认值为0
	Phone      []string `json:"phone"`      // 收信人手机号,多个手机号用逗号,分隔，每次调用最大支持2000，更多地址建议使用联系人列表功能
	Vars       string   `json:"vars"`       // 替换变量的json串
	Signature  string   `json:"signature"`  // 数字签名, 合法性验证，详情见API 验证机制
	Timestamp  string   `json:"timestamp"`  // UNIX时间戳
	Tag        string   `json:"tag"`        // 值为json 格式字符串,最大字符长度为128,比如:{"key1": "value1", "key2": "value2"}
}

func (m *Message) Default(apiUser string) {
	m.SmsUser = apiUser
}

func (m *Message) ToSend() url.Values {
	v := url.Values{}

	v.Add("smsUser", m.SmsUser)
	v.Add("templateId", convert.Int64ToString(m.TemplateID))
	v.Add("msgType", convert.Int64ToString(m.MsgType))
	v.Add("phone", strings.Join(m.Phone, ","))
	v.Add("vars", m.Vars)
	v.Add("timestamp", time.Gen13BitTimestamp()) //  fmt.Sprintf("%d", time.Now().Unix()*1000),
	v.Add("signature", m.Signature)
	return v
}

// 生成签名:
func (m *Message) Sign(smsKey string) {
	m.Signature = "" // TODO: https://www.sendcloud.net/doc/sms/#_7
}

///////////////////////////////////////////////////////////////////////////////////////////////////

// http://www.sendcloud.net/smsapi/sendVoice
type MessageVoice struct {
	SmsUser string `json:"smsUser"` //

	//
	LabelID   int64  `json:"labelId"`   // 短信标签ID
	Phone     string `json:"phone"`     // 收信人手机号
	Code      string `json:"code"`      // 验证码
	Signature string `json:"signature"` // 数字签名, 合法性验证，详情见API 验证机制
	Timestamp string `json:"timestamp"` // UNIX时间戳
	Tag       string `json:"tag"`       // 值为json 格式字符串,最大字符长度为128,比如:{"key1": "value1", "key2": "value2"}
}

func (m *MessageVoice) Default(apiUser string) {
	m.SmsUser = apiUser
}

func (m *MessageVoice) ToSend() url.Values {
	v := url.Values{}

	v.Add("smsUser", m.SmsUser)

	return v
}

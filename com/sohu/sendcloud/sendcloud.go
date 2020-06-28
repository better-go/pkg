package sendcloud

/*
SendCloud 有 国内版 和 国际版 两个版本

	- docs:
		- https://www.sendcloud.net/doc/product_email/quickin/
		- email api:
			- https://www.sendcloud.net/doc/email_v2/
	- api swagger:
		- https://www.sendcloud.net/doc/test/
*/

type SendCloud struct {
	mail *Email
	sms  *Sms
}

func NewSendCloud(opt *Option) *SendCloud {
	return &SendCloud{
		mail: NewEmail(opt.EmailApiUser, opt.EmailApiKey),
		sms:  NewSms(opt.SmsApiUser, opt.SmsApiKey),
	}
}

//
func (m *SendCloud) SendEmail() {

}

//
func (m *SendCloud) SendSms() {

}

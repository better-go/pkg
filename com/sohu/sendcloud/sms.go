package sendcloud

type Sms struct {
	smsUser string
	smsKey  string
}

func NewSms(smsUser string, smsKey string) *Sms {
	return &Sms{
		smsUser: smsUser,
		smsKey:  smsKey,
	}
}

func (m *Sms) Send() {

}

func (m *Sms) SendVoice() {

}

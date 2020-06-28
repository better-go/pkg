package sendcloud

type Sms struct {
	apiUser string
	apiKey  string
}

func NewSms(apiUser string, apiKey string) *Sms {
	return &Sms{
		apiUser: apiUser,
		apiKey:  apiKey,
	}
}

func (m *Sms) Send() {

}

func (m *Sms) SendVoice() {

}

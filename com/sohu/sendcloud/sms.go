package sendcloud

import (
	"net/http"
	"time"
)

type Sms struct {
	apiUser string
	apiKey  string

	// http:
	cli *http.Client
}

func NewSms(apiUser string, apiKey string) *Sms {
	return &Sms{
		apiUser: apiUser,
		apiKey:  apiKey,
		cli: &http.Client{
			Timeout: time.Second * 1, // 1s
		},
	}
}

func (m *Sms) Send(msg *Message) error {
	msg.Default(m.apiUser)

	// post:
	return QuickPost(m.cli, urlSms, msg.ToSend())
}

func (m *Sms) SendVoice(msg *MessageVoice) error {
	msg.Default(m.apiUser)

	// post:
	return QuickPost(m.cli, urlSmsVoice, msg.ToSend())
}

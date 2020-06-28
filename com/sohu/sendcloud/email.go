package sendcloud

import (
	"net/http"
	"time"
)

type Email struct {
	apiUser string
	apiKey  string

	// http:
	cli *http.Client
}

func NewEmail(apiUser string, apiKey string) *Email {
	return &Email{
		apiUser: apiUser,
		apiKey:  apiKey,
		cli: &http.Client{
			Timeout: time.Second * 1, // 1s
		},
	}
}

func (m *Email) Send(mail *Mail) error {
	// set default:
	mail.Default(m.apiUser, m.apiKey)

	// post:
	return QuickPost(m.cli, urlMail, mail.ToSend())
}

func (m *Email) SendTemplate(mail *MailTemplate) error {
	// set default:
	mail.Default(m.apiUser, m.apiKey)

	// post:
	return QuickPost(m.cli, urlMailTemplate, mail.ToSend())
}

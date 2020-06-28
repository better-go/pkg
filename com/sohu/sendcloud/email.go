package sendcloud

import (
	"net/http"
)

type Email struct {
	apiUser string
	apiKey  string

	cli *http.Client
}

func NewEmail(apiUser string, apiKey string) *Email {
	return &Email{
		apiUser: apiUser,
		apiKey:  apiKey,
		cli:     &http.Client{},
	}
}

func (m *Email) Send(mail *Mail) error {
	// set default:
	mail.Default(m.apiUser, m.apiKey)

	// post:
	return QuickPost(m.cli, urlMail, mail.ToSend())
}

func (m *Email) SendTemplate() {

}

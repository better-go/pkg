package email

import (
	"github.com/jordan-wright/email"
)

// option:
type Option struct {
	Identity string
	Email    string
	Password string
	Host     string
}

// content:
type MailContent email.Email

func (m *MailContent) ToMail() *email.Email {
	return &email.Email{
		ReplyTo:     m.ReplyTo,
		From:        m.From,
		To:          m.To,
		Bcc:         m.Bcc,
		Cc:          m.Cc,
		Subject:     m.Subject,
		Text:        m.Text,
		HTML:        m.HTML,
		Sender:      m.Sender,
		Headers:     m.Headers,
		Attachments: m.Attachments,
		ReadReceipt: m.ReadReceipt,
	}
}

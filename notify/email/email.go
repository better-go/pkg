package email

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

//
type Mailer struct {
	opt *Option

	e *email.Email
}

func NewMailer(opt *Option) *Mailer {
	return &Mailer{
		opt: opt,
		e:   email.NewEmail(),
	}
}

func (m *Mailer) Send(toAddr string, mail *MailContent) error {
	m.e = mail.ToMail()
	return m.e.Send(toAddr, smtp.PlainAuth(m.opt.Identity, m.opt.Email, m.opt.Password, m.opt.Host))
}

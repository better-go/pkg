package sendcloud

import "testing"

func TestSendCloud_SendEmail(t *testing.T) {
	apiUser := "test user"
	apiKey := "test key"
	//
	from := "from@gmail.com"
	to := "to@gmail.com"

	// new:
	mailer := NewEmail(apiUser, apiKey)

	// send:
	err := mailer.Send(&Mail{
		ApiUser:         apiUser,
		ApiKey:          apiKey,
		AddrFrom:        from,
		AddrTo:          []string{to,},
		AddrCc:          nil,
		AddrBcc:         nil,
		FromName:        "",
		ReplyTo:         nil,
		Subject:         "test email",
		Html:            "verify code 2233",
		Plain:           "",
		Headers:         "",
		Attachments:     nil,
		UseNotification: false,
		UseAddressList:  false,
	})

	t.Log("resp:", err)

}

func TestSendCloud_SendSms(t *testing.T) {

}

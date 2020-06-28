package sendcloud

import "testing"

func TestSendCloud_SendEmail(t *testing.T) {
	apiUser := "test_user" // test_user
	apiKey := "test_key"   // test_key
	//
	from := "from@gmail.com" // from
	to := "to@gmail.com"     // to

	// new:
	mailer := NewEmail(apiUser, apiKey)

	// send:
	err := mailer.Send(&Mail{
		//ApiUser:         apiUser,
		//ApiKey:          apiKey,
		AddrFrom: from,
		AddrTo:   []string{to},
		AddrCc:   nil,
		AddrBcc:  nil,
		FromName: "",
		ReplyTo:  nil,
		Subject:  "test email",
		Html: `
			<h1>your verify code</h1>: <div style="color:blue">3322</div>
		`,
		Plain:           "",
		Headers:         "",
		Attachments:     nil,
		UseNotification: false,
		UseAddressList:  false,
	})

	/*
		=== RUN   TestSendCloud_SendEmail
		2020-06-29T04:48:04.081+0800	DEBUG	log/log.go:78	sendcould post resp:
			{
				"result":true,"statusCode":200,
				"message":"请求成功",
				"info":{
					"emailIdList":["1593377284034_144383_11771_2575.sc-10_9_13_218-inbound0$xxuser@outlook.com"]
			}},
			err: <nil>
		--- PASS: TestSendCloud_SendEmail (0.19s)
	*/
	t.Log("send email resp:", err)

}

func TestSendCloud_SendSms(t *testing.T) {
	smsUser := "test_user" // test_user
	smsKey := ""           // for gen Signature
	//
	phone := "test_phone_no" // phone_no
	vars := ""

	// new:
	smsSender := NewSms(smsUser, smsKey)

	// send:
	err := smsSender.Send(&Message{
		SmsUser:    smsUser,
		TemplateID: 0,
		LabelID:    0,
		MsgType:    0,
		Phone:      []string{phone},
		Vars:       vars,
		Signature:  "",
		Timestamp:  "",
		Tag:        "",
	})

	t.Log("send sms resp:", err)
}

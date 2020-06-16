package sendcloud

type Email struct {
	apiUser string
	apiKey  string
}

func NewEmail(apiUser string, apiKey string) *Email {
	return &Email{
		apiUser: apiUser,
		apiKey:  apiKey,
	}
}

func (m *Email) Send() {

}

func (m *Email) SendTemplate() {

}

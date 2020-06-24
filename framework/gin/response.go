package gin

type ResponseData struct {
	Code    int64
	Message string
	Data    interface{}
}

func (m *ResponseData) From(code int64, message string, data interface{}) {
	m.Code = code
	m.Message = message
	m.Data = data
}

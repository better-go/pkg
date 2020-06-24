package errors

import "testing"

func TestHttpError_ErrCode(t *testing.T) {
	err1 := HttpError{
		Code:   200,
		Id:     "ok",
		Detail: "this is ok.",
	}

	t.Log(err1.GetCode(), err1.GetId(), err1.GetDetail())
}

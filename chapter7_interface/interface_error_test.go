package chapter7_interface

import (
	"fmt"
	"testing"
)

type BizError struct {
	code int
	msg  string
}

func (e *BizError) Error() string {
	return fmt.Sprintf("errCode:%v errMsg:%v", e.code, e.msg)
}

func NewBizError(code int, msg string) *BizError {
	return &BizError{code: code, msg: msg}
}

func Http(code int) error {
	if code >= 500 {
		return NewBizError(500, "server error")
	}
	return nil
}

func TestHttp(t *testing.T) {
	err := Http(500)
	if nil != err {
		fmt.Println(err)
	}
}

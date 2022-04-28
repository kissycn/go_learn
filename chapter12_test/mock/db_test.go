package mock

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	// 当 Get() 的参数为 Tom，则返回 100 error 为nil ，这称之为打桩(stub)
	m.EXPECT().Get(gomock.Eq("tom")).Return(100, nil)

	if v := GetFromDB(m, "tom"); v == -1 {
		t.Fatal("expected -1, but got", v)
	} else {
		fmt.Println("res=", v)
	}
}

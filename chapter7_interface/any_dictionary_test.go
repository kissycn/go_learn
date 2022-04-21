package chapter7_interface

import (
	"fmt"
	"testing"
)

type any interface{}

type Dic struct {
	data map[any]any
}

func (d *Dic) Set(key, value any) {
	d.data[key] = value
}

func (d *Dic) Get(key any) any {
	return d.data[key]
}

func (d *Dic) Visit(callback func(k, v any)) {
	for k, v := range d.data {
		callback(k, v)
	}
}

func (d *Dic) Clear() {
	d.data = make(map[any]any)
}

func NewDic() *Dic {
	d := new(Dic)
	d.Clear()
	return d
}

func TestDic(t *testing.T) {
	var dic = NewDic()
	dic.Set("k1", "v1")
	dic.Set("k2", "v2")

	fmt.Println(dic.Get("k1"))

	dic.Visit(func(k, v any) {
		fmt.Printf("key:%s val:%s \n", k, v)
	})
}

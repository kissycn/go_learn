package chapter11_file

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Info struct {
	Name   string
	Url    string
	Course []string
}

func TestReadJson(t *testing.T) {
	jfile, err := os.Open("./info.json")
	if nil != err {
		fmt.Println(err)
	}

	defer jfile.Close()

	var info []Info
	decoder := json.NewDecoder(jfile)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", info)
	}
}

package chapter7_interface

import (
	"fmt"
	"testing"
)

type File struct {
}

type DataWriter interface {
	WriteData(data interface{}) error
}

func (f *File) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data, " by file")
	return nil
}

type Socket struct {
}

func (s *Socket) WriteData(data interface{}) error {
	fmt.Println("WriteData:", data, " by socket")
	return nil
}

func TestDataWriter(t *testing.T) {
	//file := new(File)
	var writer DataWriter = new(File)
	writer.WriteData("test")

	var writer1 DataWriter = new(Socket)
	writer1.WriteData("test")
}

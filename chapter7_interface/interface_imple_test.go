package chapter7_interface

import (
	"fmt"
	"testing"
)

//  ************** 一个类型实现多个接口  **************
type Writer interface {
	WriteDate(s string) error
}

type Closer interface {
	Close()
}

type Socket1 struct {
}

func (s1 *Socket1) WriteDate(s string) error {
	fmt.Println("write by socket", s)
	return nil
}

func (s *Socket1) Close() {
	fmt.Println("socket close")
}

func useWrite(w Writer, str string) {
	w.WriteDate(str)
}

func useClose(closer Closer) {
	closer.Close()
}

func TestSocket1(t *testing.T) {
	var socket = new(Socket1)
	useWrite(socket, "test")
	useClose(socket)
}

//  ************** 一个接口多个类型实现  **************

type Service interface {
	Start()
	Log()
}

type Logger1 struct {
}

func (log *Logger1) Log() {
	fmt.Println("system log")
}

type GameService struct {
	Logger1
}

func (gs *GameService) Start() {
	fmt.Println("game start")
}

func TestGameService(t *testing.T) {
	var gs = new(GameService)
	gs.Start()
	gs.Log()
}

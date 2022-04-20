package chapter6_struct

import (
	"fmt"
	"testing"
)

type Color struct {
	R, G, B byte
}

type Player struct {
	Name        string
	HealthPoint int
	MagicPoint  int
}

func TestStructDeclareByNew(t *testing.T) {
	var kk = new(Player)
	kk.Name = "cc"
	kk.HealthPoint = 100
	kk.MagicPoint = 99

	fmt.Println(kk)
}

type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令变量
	Comment string // 指令备注
}

func TestCommand(t *testing.T) {
	var i int = 1
	c := &Command{
		Name:    "ping",
		Var:     &i,
		Comment: "icmp",
	}
	fmt.Println(c)

	c1 := new(Command)
	c1.Name = "telnet"
	c1.Var = &i
	c1.Comment = "telnet command"
	fmt.Println(c1)

	c2 := Command{
		Name:    "ls",
		Comment: "list direction",
	}

	c2.Name = "11"
	fmt.Println(c2)

	var c3 Command
	c3.Name = "ps"
}

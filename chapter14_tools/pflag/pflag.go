package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

// 方式一：通过赋值将 flag 与变量绑定，格式：参数名，默认值，参数说明
var ipAddr = pflag.String("ipAddr", "127.0.0.1", "bind server address.")

// 方式二：通过 Var() 函数将 flag 与变量绑定
var port int
var name string
var level string

func initFlag() {
	// 格式：参数绑定变量名，参数名，默认值，参数说明
	pflag.IntVar(&port, "port", 8080, "listener port.")
	// 格式：参数绑定变量名，参数简写，参数名，默认值，参数说明
	pflag.StringVarP(&name, "name", "n", "admin", "账户名")
	//
	pflag.StringVarP(&level, "level", "l", "debug", "log level")
	pflag.CommandLine.MarkDeprecated("level", "please use --log-level instead")
}

func main() {
	initFlag()
	pflag.Parse()

	fmt.Println("server:", *ipAddr)
	fmt.Println("port:", port)
	fmt.Println("name:", name)
}

package chapter5_func

import (
	"fmt"
	"strings"
	"testing"
)

func TestFuncDeclare(t *testing.T) {

}

func f1(x, y int) int {
	return x + y
}

func f2(x, y int) (sum int) {
	sum = x + y
	return
}

func f3() (x, y int) {
	x = 1
	y = 2
	return
}

func f4() (x, y int) {
	x = 1
	return x, 1
}

/**
  函数返回值---示例小程序
*/
const (
	SecondPerMinute = 60

	SecondPerHour = SecondPerMinute * 60

	SecondPerDay = SecondPerHour * 60
)

func ParserDate(seconds int) (minute, hour, day int) {
	minute = seconds / SecondPerMinute

	hour = seconds / SecondPerHour

	day = seconds / SecondPerDay

	return
}

func TestDateParser(t *testing.T) {
	fmt.Println(ParserDate(18000))
}

/**
  函数链--示例小程序
*/

func StrProcess(str []string, funcChain []func(string) string) {
	for index, value := range str {
		result := value
		for _, f := range funcChain {
			result = f(result)
		}
		str[index] = result
	}
}

func removePrefix(s string) string {
	return strings.TrimPrefix(s, "go")
}

func TestFuncChain(t *testing.T) {
	var list3 = make([]string, 1, 2)
	list3 = append(list3, "111", "222")
	fmt.Println(list3)

	list := []string{
		"go scanner",
		"go compiler",
		"go run",
	}

	funcList := []func(string) string{
		removePrefix,
		strings.TrimSpace,
		strings.ToUpper,
		func(s string) string {
			return s + ";"
		},
	}

	StrProcess(list, funcList)
	for _, str := range list {
		fmt.Println(str)
	}
}

/**

 */
func TestVisit(t *testing.T) {
	list := []string{
		"one",
		"two",
		"three",
	}

	visit(list, func(s string) {
		fmt.Println(s)
	})
}
func visit(strings []string, f func(string)) {
	for _, s := range strings {
		f(s)
	}
}

/**
  可变参数
*/
func variableArgs(f func(string, interface{}), args ...interface{}) {
	for _, value := range args {
		switch value.(type) {
		case int:
			f("int", value)
		case string:
			f("string", value)
		default:
			fmt.Println("default")
		}
	}
}

func TestVariable(t *testing.T) {
	variableArgs(func(s string, i interface{}) {
		fmt.Printf("type:%v value:%v \n", s, i)
	}, 1, "one")
}

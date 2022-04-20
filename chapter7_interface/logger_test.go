package chapter7_interface

type LogWriter interface {
	Write(log interface{}) error
}

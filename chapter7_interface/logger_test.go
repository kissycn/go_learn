package chapter7_interface

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

type LogWriter interface {
	Write(date interface{}) error
}

/**
  Log 管理
*/
type Logger struct {
	Writers []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	if nil != writer {
		l.Writers = append(l.Writers, writer)
	}
}

func (l *Logger) Log(data interface{}) {
	for _, w := range l.Writers {
		w.Write(data)
	}
}

func NewLogger() *Logger {
	return new(Logger)
}

type FileWriter struct {
	file *os.File
}

func (fw *FileWriter) SetFile(fn string) (err error) {
	if fw.file != nil {
		fw.file.Close()
	}

	fw.file, err = os.Create(fn)
	return err
}

func (fw *FileWriter) Write(data interface{}) error {
	// 日志文件可能没有创建成功
	if fw.file == nil {
		// 日志文件没有准备好
		return errors.New("file not created")
	}
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入文件中
	_, err := fw.file.Write([]byte(str))
	return err
}

func NewFileWriter() *FileWriter {
	return new(FileWriter)
}

type ConsoleWriter struct {
}

func (cw *ConsoleWriter) Write(data interface{}) error {
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWriter() *ConsoleWriter {
	return new(ConsoleWriter)
}

func TestLogger(t *testing.T) {
	var logger = NewLogger()

	var fileLogger = NewFileWriter()
	if err := fileLogger.SetFile("log.log"); err != nil {
		fmt.Println(err)
	} else {
		logger.RegisterWriter(fileLogger)
	}

	var consoleLog = NewConsoleWriter()
	logger.RegisterWriter(consoleLog)

	logger.Log("Hello World")
}

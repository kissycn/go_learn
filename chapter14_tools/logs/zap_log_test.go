package logs

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func TestZapLoggerNew(t *testing.T) {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()
	url := "http://marmotedu.com"

	logger.Info(
		"failed to fetch URL by logger",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL by sugar logger",
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func TestZapLoggerEnv(t *testing.T) {
	log1 := zap.NewExample()
	log2, _ := zap.NewDevelopment()
	log3, _ := zap.NewProduction()
	log1.Debug("This is a DEBUG message for example.")
	log1.Info("This is an INFO message for example.")

	log2.Debug("This is a DEBUG message for dev.")
	log2.Info("This is an INFO message for dev.")

	log3.Debug("This is a DEBUG message for pro.")
	log3.Info("This is an INFO message for pro.")
}

func TestZapAddFields(t *testing.T) {
	logger := zap.NewExample(zap.Fields(
		zap.Int("userID", 10),
		zap.String("requestID", "fbf54504"),
	))

	logger.Debug("This is a debug message")
	logger.Info("This is a info message")
}

func TestZapNamespace(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.Info("tracked some metrics",
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)

	logger2 := logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	)
	logger2.Info("tracked some metrics")
}

func TestZapAddCaller(t *testing.T) {
	logger, _ := zap.NewProduction(zap.AddCaller())
	defer logger.Sync()

	logger.Info("hello world")
}

func TestZapAddStackTrace(t *testing.T) {
	logger, _ := zap.NewProduction(zap.AddStacktrace(zapcore.WarnLevel))
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	f1()
}

func f1() {
	zap.L().Warn("f1() caller")
	f2("world")
}

func f2(word string) {
	fmt.Println("hello", word)
	zap.L().Warn("f1() caller")
}

func TestZapGlobal(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
	f1()
}

func TestZapRotation(t *testing.T) {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	fileWriteSync := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./info.log",
		MaxSize:    1,     // 单位：M
		MaxBackups: 10,    // 最大保留日志文件数量
		MaxAge:     28,    // 日志文件保留天数
		Compress:   false, //是否压缩处理
	})

	fileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(fileWriteSync,
			zapcore.AddSync(os.Stdout)), zapcore.InfoLevel)

	logger := zap.New(fileCore, zap.AddCaller()) //AddCaller()为显示文件名和行号
	logger.Info("hello world")
	logger.Error("hello world")
}

func TestZapLevelFile(t *testing.T) {
	var coreArr []zapcore.Core
	encoderConfig := zap.NewProductionEncoderConfig()
	//指定时间格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		//error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		//info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/info.log", //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    1,                //文件大小限制,单位MB
		MaxBackups: 5,                //最大保留日志文件数量
		MaxAge:     30,               //日志文件保留天数
		Compress:   false,            //是否压缩处理
	})
	//error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "./log/error.log", //日志文件存放目录
		MaxSize:    1,                 //文件大小限制,单位MB
		MaxBackups: 5,                 //最大保留日志文件数量
		MaxAge:     30,                //日志文件保留天数
		Compress:   false,             //是否压缩处理
	})
	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	infoFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(
			infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)

	//第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	errorFileCore := zapcore.NewCore(encoder,
		zapcore.NewMultiWriteSyncer(
			errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	//zap.AddCaller()为显示文件名和行号，可省略
	logger := zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
	logger.Info("hello info")
	logger.Debug("hello debug")
	logger.Error("hello error")
}

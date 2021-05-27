package zap

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// 	zaprotatelogs "github.com/lestrrat-go/file-rotatelogs"

var ZapLogger *zap.Logger

func InitLogger() {
	//zap.Config{}
	// 此处的配置是从我的项目配置文件读取的，读者可以根据自己的情况来设置
	//zap := Zap{
	//	Level:         "",
	//	Format:        "",
	//	Prefix:        "",
	//	Director:      "",
	//	LinkName:      "",
	//	ShowLine:      false,
	//	EncodeLevel:   "",
	//	StacktraceKey: "",
	//	LogInConsole:  false,
	//}
	logPath := "1.log"
	name := "Name"
	debug := true

	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    128,     // 每个日志文件保存的大小 单位:M
		MaxAge:     7,       // 文件最多保存多少天
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别， 默认是 zap.InfoLevel
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是开发环境，同时在控制台上也输出
	if debug {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		//zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	field := zap.Fields(zap.String("appName", name))

	// 构造日志
	ZapLogger = zap.New(core, caller, development, field)
	ZapLogger.Info("log 初始化成功")
	ZapLogger.Debug("log debug")
	ZapLogger.Error("log error")

}

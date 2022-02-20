package boot

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/app/global"
	"main/utils"
	"os"
	"time"
)

// 格式化时间
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

var f = false

// 编码器设置
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",                     // 日志消息键
		LevelKey:       "level",                       // 日志等级键
		TimeKey:        "time",                        // 时间键
		NameKey:        "logger",                      // 日志记录器名
		CallerKey:      "caller",                      // 日志文件信息键
		StacktraceKey:  "stacktrace",                  // 堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // 友好日志等级名大小写（info INFO）
		EncodeTime:     CustomTimeEncoder,             // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.FullCallerEncoder,     // 日志文件信息 short（包/文件.go:行号） full (文件位置.go:行号)
	}
	if f == false {
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		f = true
	} else {
		switch {
		case g.Config.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
			config.EncodeLevel = zapcore.LowercaseLevelEncoder
		case g.Config.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
			config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		case g.Config.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
			config.EncodeLevel = zapcore.CapitalLevelEncoder
		case g.Config.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
			config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		default:
			config.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	}
	if g.Config.Zap.EncodeCaller == "ShortCallerEncoder" {
		config.EncodeCaller = zapcore.ShortCallerEncoder
	}
	return config
}

// 读写器设置
func getWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    1,    // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 100,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	if g.Config.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

func ZapSetup() {
	// 判断是否有Director文件夹
	if err := utils.IsNotExistMkDir(g.Config.Zap.Director); err != nil {
		fmt.Printf("Create %v directory\n", g.Config.Zap.Director)
		_ = os.Mkdir(g.Config.Zap.Director, os.ModePerm)
	}
	dynamicLevel := zap.NewAtomicLevel()
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(getEncoder(), os.Stdout, dynamicLevel), //控制台输出
		//日志文件输出,按等级归档
		zapcore.NewCore(getEncoder(), getWriteSyncer(fmt.Sprintf("./%s/all/server_all.log", g.Config.Zap.Director)), zapcore.DebugLevel),
		zapcore.NewCore(getEncoder(), getWriteSyncer(fmt.Sprintf("./%s/debug/server_debug.log", g.Config.Zap.Director)), debugPriority),
		zapcore.NewCore(getEncoder(), getWriteSyncer(fmt.Sprintf("./%s/info/server_info.log", g.Config.Zap.Director)), infoPriority),
		zapcore.NewCore(getEncoder(), getWriteSyncer(fmt.Sprintf("./%s/warn/server_warn.log", g.Config.Zap.Director)), warnPriority),
		zapcore.NewCore(getEncoder(), getWriteSyncer(fmt.Sprintf("./%s/error/server_error.log", g.Config.Zap.Director)), errorPriority),
	}
	logger := zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer logger.Sync()
	// 将当前日志等级设置为Debug
	// 注意日志等级低于设置的等级，日志文件也不分记录
	dynamicLevel.SetLevel(zap.DebugLevel)
	sugar := logger.Sugar()
	//设置全局logger
	g.Logger = sugar
	g.Logger.Info("Initialize logger successfully!")
	//sugar.Debug("test")
	//sugar.Warn("test")
	//sugar.Error("test")
	//sugar.DPanic("test")
	//sugar.Panic("test") //打印后程序停止,defer执行
	//sugar.Fatal("test") //打印后程序停止,defer不执行
}

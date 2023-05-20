package logger

import (
	"gin-web-scaffolding/settings"
	"os"

	"go.uber.org/zap"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap/zapcore"
)

// @file      : logger.go
// @author    : 八宝糖
// @contact   : 1013269096@qq.com
// -------------------------------------------

// Init 初始化日志
func Init(cfg *settings.LogConfig, mode string) (err error) {
	writerSyncer := getLogWriter(
		cfg.FileName,
		cfg.MaxSize,
		cfg.MaxBackups,
		cfg.MaxAge)
	encoder := getEncoder()
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	var core zapcore.Core
	if mode == "dev" {
		// 开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// 将日志处理记录到文件中，同事打印到终端上
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writerSyncer, l),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writerSyncer, l)
	}
	lg := zap.New(core, zap.AddCaller())
	//  替换zap包中全局的logger实例，后续在其他包中只需要使用zap.L()调用即可
	zap.ReplaceGlobals(lg)
	return
}

// getLogWriter  定制自己的log writer，作用是对日志进行分割
func getLogWriter(filename string, maxSize, maxBackups, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// getEncoder 定制Encoder
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

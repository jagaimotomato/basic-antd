package logger

import (
	"basic-antd/init/config"
	"basic-antd/init/global"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Setup() (err error) {
	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   config.LoggerConfig.Filename,
		MaxSize:    config.LoggerConfig.MaxSize,
		MaxAge:     config.LoggerConfig.MaxAge,
		MaxBackups: config.LoggerConfig.MaxBackups,
	})
	encoder := zapcore.NewJSONEncoder(getEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(getEncoderConfig())
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(config.LoggerConfig.Level))
	if err != nil {
		return
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, writeSyncer, l),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), l),
	)

	global.Lg = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(global.Lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	return
}

func getEncoderConfig() zapcore.EncoderConfig {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return encoderConfig
}

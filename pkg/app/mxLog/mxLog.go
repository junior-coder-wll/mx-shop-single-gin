//@Author: wulinlin
//@Description:
//@File:  mxLog
//@Version: 1.0.0
//@Date: 2023/03/24 14:02

package mxLog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// 自定义TimeEncoder，使时间输出格式更加清晰
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

type LogConfig struct {
	Level       string `mapstructure:"level"`
	Path        string `mapstructure:"path"`
	Filename    string `mapstructure:"filename"`
	ErrFilename string `mapstructure:"err_filename"`
	MaxAge      int    `mapstructure:"max_size"`
	MaxSize     int    `mapstructure:"max_age"`
	MaxBackups  int    `mapstructure:"max_backups"`
	Format      string `mapstructure:"format"`
	Compress    bool   `mapstructure:"compress"`
}

func NewLogger(isProd bool, logCfg *LogConfig) (*zap.Logger, error) {
	// 1. 设置日志级别
	level := zap.NewAtomicLevel()

	// 2. 配置编码器
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder, // 使用大写的彩色字体来显示级别，例如INFO显示为蓝色，WARN显示为黄色，ERROR显示为红色。EncodeTime用来设置时间输出格式
		EncodeTime:   TimeEncoder,                      // 自定义时间输出格式
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	// 设置日志输出的路径
	var cores []zapcore.Core
	if !isProd {
		// 开发环境，日志输出到控制台，采用text格式
		consoleCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		)
		cores = append(cores, consoleCore)
		// 构造Logger
		logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

		// 替换zap包中全局的logger实例，之后在项目中使用 zap.L().xxx即可使用日志
		zap.ReplaceGlobals(logger)

		return logger, nil
	}
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	// 生产环境，日志输出到文件，采用json格式
	// 一个文件用于记录info及以上级别的日志
	infoWrite := lumberjack.Logger{
		Filename:   logCfg.Filename,   // 日志文件路径
		MaxSize:    logCfg.MaxSize,    // 每个日志文件的最大大小，单位：MB
		MaxBackups: logCfg.MaxBackups, // 日志文件最多保存的个数
		MaxAge:     logCfg.MaxAge,     // 文件最多保存多少天
		Compress:   logCfg.Compress,   // 是否压缩
	}
	// JSON格式输出
	infoEncoder := zapcore.NewJSONEncoder(encoderCfg)
	infoCore := zapcore.NewCore(infoEncoder, zapcore.AddSync(&infoWrite), level)

	// 一个文件用于记录error及以上级别的日志
	errorWriter := lumberjack.Logger{
		Filename:   logCfg.ErrFilename, // 日志文件路径
		MaxSize:    logCfg.MaxSize,     // 每个日志文件的最大大小，单位：MB
		MaxBackups: logCfg.MaxBackups,  // 日志文件最多保存的个数
		MaxAge:     logCfg.MaxAge,      // 文件最多保存多少天
		Compress:   logCfg.Compress,    // 是否压缩
	}
	errEncoder := zapcore.NewJSONEncoder(encoderCfg)
	errLevel := zap.NewAtomicLevel()
	errLevel.SetLevel(zapcore.ErrorLevel)
	errCore := zapcore.NewCore(errEncoder, zapcore.AddSync(&errorWriter), errLevel)

	cores = append(cores, infoCore, errCore)

	// 构造Logger
	logger := zap.New(zapcore.NewTee(cores...), zap.AddCaller())

	// 替换zap包中全局的logger实例，之后在项目中使用 zap.L().xxx即可使用日志
	zap.ReplaceGlobals(logger)

	return logger, nil

}

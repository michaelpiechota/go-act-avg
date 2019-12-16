package actions

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var cfg = zap.Config{
	Encoding:         "json",
	Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
	OutputPaths:      []string{"stderr"},
	ErrorOutputPaths: []string{"stderr"},
	EncoderConfig: zapcore.EncoderConfig{
		MessageKey: "message",

		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,

		TimeKey:    "time",
		EncodeTime: zapcore.ISO8601TimeEncoder,

		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
	},
}

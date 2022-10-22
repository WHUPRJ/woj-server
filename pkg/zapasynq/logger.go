package zapasynq

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func New(zapLogger *zap.Logger) Logger {
	return Logger{
		logger: zapLogger,
	}
}

func (l Logger) Debug(args ...interface{}) { l.logger.Sugar().Debugf(args[0].(string), args[1:]...) }

func (l Logger) Info(args ...interface{}) { l.logger.Sugar().Infof(args[0].(string), args[1:]...) }

func (l Logger) Warn(args ...interface{}) { l.logger.Sugar().Warnf(args[0].(string), args[1:]...) }

func (l Logger) Error(args ...interface{}) { l.logger.Sugar().Errorf(args[0].(string), args[1:]...) }

func (l Logger) Fatal(args ...interface{}) { l.logger.Sugar().Fatalf(args[0].(string), args[1:]...) }

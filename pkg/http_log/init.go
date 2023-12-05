package http_log

import (
	"errors"
	"knowledge-tree/pkg/program"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type output struct {
	channel chan struct{} //阻塞管道
	target  int8          //日志输出目标
	value   int           //日志分割值(时间分割模式:时间值,大小分割模式:文件大小值)
	dir     string        //日志文件目录名
	logger  *zap.SugaredLogger
}

type HttpLog interface {
	Print(string, ...any)
	Printf(string, ...any)
}

// InitHttpLogger 初始化http服务日志记录器
func InitHttpLogger(opt ...Options) (HttpLog, error) {
	var l = new(output)
	l.channel = make(chan struct{}, 1)
	for _, v := range opt {
		v.apply(l)
	}
	err := l.initLogger()
	if err != nil {
		return nil, err
	}
	return l, nil
}

// 初始化日志
func (l *output) initLogger() error {
	l.value = time.Now().Day()
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = ""
	config.LevelKey = ""
	config.CallerKey = ""
	encoder := zapcore.NewConsoleEncoder(config)
	switch l.target {
	case outputFile:
		err := os.MkdirAll(l.dir, os.ModePerm)
		if err != nil {
			return err
		}
		name, err := program.GetProgramName()
		if err != nil {
			return err
		}
		filename := l.dir + "/" + name + "_" +
			time.Now().Format(time.DateOnly) + ".log"
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}
		syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(file))
		l.logger = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		return nil
	case outputTerminal:
		syncer := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
		l.logger = zap.New(
			zapcore.NewCore(encoder, syncer, zapcore.DebugLevel),
			zap.AddCaller(),      //打印行号
			zap.AddCallerSkip(1), //向上跳一层
		).Sugar()
		return nil
	default:
		return errors.New("unknown log output target... ")
	}
}

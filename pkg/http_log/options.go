package http_log

type Options interface {
	apply(*output)
}

type outputFunction func(*output)

func (o outputFunction) apply(l *output) {
	o(l)
}

// 日志输出目标
const (
	_              int8 = iota //无输出模式(默认终端)
	outputFile                 //输出到文件
	outputTerminal             //输出到终端
	//outputFileAndTerminal             //输出到文件和终端
)

// LogOutputToFile 日志输出到文件
func LogOutputToFile(dir string) Options {
	return outputFunction(func(l *output) {
		l.target = outputFile
		l.dir = dir
	})
}

// LogOutputToTerminal 日志输出到终端
func LogOutputToTerminal() Options {
	return outputFunction(func(l *output) {
		l.target = outputTerminal
	})
}

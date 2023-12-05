package logs

import (
	"log"
	"runtime/debug"
)

// RecoverPrint panic打印
func RecoverPrint() {
	err := recover()
	if err != nil {
		return
	}
	log.Println("[panic]\n", err)
	debug.PrintStack()
}

// RecoverPrintf 自定义panic打印内容
func RecoverPrintf(desc string, args ...any) {
	err := recover()
	if err == nil {
		return
	}
	if len(args) > 0 {
		log.Printf(desc+"\n%v\n", append(args, err)...)
		debug.PrintStack()
		return
	}
	log.Printf(desc+"\n%v\n", err)
	debug.PrintStack()
}

// RecoverLogPrint panic打印 自定义输出位置
func RecoverLogPrint(logger Logger) {
	err := recover()
	if err == nil {
		return
	}
	logger.StackDPanic("[panic]\n", err)
}

// RecoverLogPrintf 自定义panic打印内容 自定义输出位置
func RecoverLogPrintf(logger Logger, desc string, args ...any) {
	err := recover()
	if err == nil {
		return
	}
	if len(args) > 0 {
		logger.StackDPanicf(desc+"\n%v", append(args, err)...)
		return
	}
	logger.StackDPanicf(desc+"\n%v", err)
}

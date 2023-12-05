package http_log

import (
	"log"
	"time"
)

func (l *output) Print(str string, args ...any) {
	l.check()
	l.logger.Debugf(str, args...)
}

func (l *output) Printf(str string, args ...any) {
	l.check()
	l.logger.Debugf(str, args...)
}

var last bool

// 校验
func (l *output) check() {
	l.channel <- struct{}{}
	if l.checkOption() {
		err := l.initLogger()
		if err != nil && !last {
			log.Println(err)
			last = true
		}
	}
	<-l.channel
}

// 校验是否要重新创建日志记录
func (l *output) checkOption() bool {
	if l.target == outputTerminal {
		return false
	}
	if l.value != time.Now().Day() {
		return true
	}
	return false
}

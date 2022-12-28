package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

// 克隆日志实例
func (l *Logger) clone() *Logger {
	nl := *l
	return &nl
}

// 设置日志等级
func (l *Logger) WithLevel(level Level) *Logger {
	ll := l.clone()
	ll.level = level
	return ll
}

// 设置日志公共字段
func (l *Logger) WithFields(fields Fields) *Logger {
	ll := l.clone()
	if ll.fields == nil {
		ll.fields = make(Fields)
	}
	for k, v := range fields {
		ll.fields[k] = v
	}
	return ll
}

// 设置日志上下文
func (l *Logger) WithContext(ctx context.Context) *Logger {
	ll := l.clone()
	ll.ctx = ctx
	return ll
}

// 设置某一层调用栈信息
func (l *Logger) WithCaller(skip int) *Logger {
	ll := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		ll.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return ll
}

// 设置整个调用栈信息
func (l *Logger) WithCallersFrames() *Logger {
	maxCallersDepth := 25
	minCallersDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallersDepth)
	depth := runtime.Callers(minCallersDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	ll := l.clone()
	ll.callers = callers
	return ll
}

// 日志格式化
func (l *Logger) JSONFormat(msg string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = l.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	return data
}

// 日志输出
func (l *Logger) Output(msg string) {
	body, _ := json.Marshal(l.JSONFormat(msg))
	content := string(body)
	switch l.level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// 日志分级输出
func (l *Logger) Debug(v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.WithLevel(LevelDebug).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.WithLevel(LevelInfo).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.WithLevel(LevelWarn).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.WithLevel(LevelError).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprint(v...))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.WithLevel(LevelFatal).Output(fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprint(v...))
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.WithLevel(LevelPanic).Output(fmt.Sprintf(format, v...))
}

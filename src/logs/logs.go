package logs

import (
	"fmt"
	"gif_generator/src/config"
	"go.uber.org/fx"
	"log"
	"os"
	"time"
)

// ConsoleLogger 是 Logger 接口的实现，输出日志到控制台
type ConsoleLogger struct {
	logLevel int // 用于控制输出的日志级别
}

// 日志等级常量
const (
	DebugLevel = iota // 0
	InfoLevel         // 1
	WarnLevel         // 2
	ErrorLevel        // 3
)

// Params 注入参数
type Params struct {
	fx.In
	Config *config.GConfig
}

// NewConsoleLogger 创建并返回一个 ConsoleLogger 实例
func NewConsoleLogger(p Params) *ConsoleLogger {
	return &ConsoleLogger{
		logLevel: p.Config.LogLevel,
	}
}

func (l *ConsoleLogger) log(level int, levelStr, message string) {
	if level >= l.logLevel {
		// 获取当前时间戳
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// 构建日志内容
		logMessage := fmt.Sprintf("[%s] [%s]: %s\n", timestamp, levelStr, message)
		log.Println(logMessage)

		// 打开或创建日志文件 (根目录下的log.txt)
		file, err := os.OpenFile(
			fmt.Sprintf("./logs/%s.txt", time.Now().Format("2006-01-02")),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening log file: %v\n", err)
			return
		}
		defer file.Close()

		// 写入日志文件
		if _, err := file.WriteString(logMessage); err != nil {
			fmt.Printf("Error writing to log file: %v\n", err)
		}
	}
}

// Debug 输出调试信息
func (l *ConsoleLogger) Debug(message string) {
	l.log(DebugLevel, "DEBUG", message)
}

// Info 输出普通信息
func (l *ConsoleLogger) Info(message string) {
	l.log(InfoLevel, "INFO", message)
}

// Warn 输出警告信息
func (l *ConsoleLogger) Warn(message string) {
	l.log(WarnLevel, "WARN", message)
}

// Error 输出错误信息
func (l *ConsoleLogger) Error(message string) {
	l.log(ErrorLevel, "ERROR", message)
}

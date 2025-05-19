package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// Log 创建一个新的日志记录器
var Log = logrus.New()

func init() {
	// 配置日志文件的切割
	//logFile := &lumberjack.Logger{
	//	Filename:   getLogFileName(), // 根据日期动态生成日志文件名
	//	MaxSize:    10,               // 每个日志文件最大 10MB
	//	MaxBackups: 30,               // 保留的旧日志文件最多30个
	//	MaxAge:     7,                // 日志文件最多保存7天
	//	Compress:   true,             // 是否压缩日志
	//}
	// 代码具体位置
	//Log.SetReportCaller(true)
	// 设置日志格式和输出
	Log.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true, // 默认格式的 等级的颜色
		DisableColors:    false,
		ForceQuote:       true,
		DisableQuote:     true, // 类似[0001],
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableSorting:   false,
	})
	// 创建一个多写入器，将日志同时写入控制台和文件
	//multiWriter := io.MultiWriter(os.Stdout, logFile)

	//Log.SetOutput(multiWriter) // 多写入器

	// 设置日志级别
	Log.SetLevel(logrus.InfoLevel)

	Log.Info("Logger initialized")
}

// 获取每天的日志文件名
func getLogFileName() string {
	currentTime := time.Now()
	logDir := "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, 0755)
		if err != nil {
			return ""
		}
	}
	return fmt.Sprintf("%s/%s.log", logDir, currentTime.Format("2006-01-02"))
}

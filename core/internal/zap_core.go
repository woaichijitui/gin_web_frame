package internal

import (
	"fmt"
	"gin_web_frame/global"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"time"
)

// 定义日志前缀
func SetLogPrefix() {
	if global.CONFIG.Zap.Prefix != "" {
		LogPrefix = "[" + global.CONFIG.Zap.Prefix + "]"
		return
	}
	if global.CONFIG.System.ServerName != "" {
		LogPrefix = "[" + global.CONFIG.System.ServerName + "]"
	}

}

var (
	LogPrefix = "myapp"
)

// MyEncoder 自定义 Encoder 结构体
type MyEncoder struct {
	zapcore.Encoder
	ErrFile    *os.File
	File       *os.File
	CurrentDay string
	LogRootDir string
}

// EncodeEntry 方法实现了对日志条目的自定义编码
func (e *MyEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (buf *buffer.Buffer, err error) {
	buf, err = e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	// 在日志行的最前面添加前缀
	data := buf.String()
	buf.Reset()
	buf.AppendString(LogPrefix + "  " + data)

	// 检查是否需要写入新的文件中
	currentDay := time.Now().Format("2006-01-02")
	logDir := filepath.Join(e.LogRootDir, "/", currentDay)
	if currentDay != e.CurrentDay {
		if e.File != nil {
			e.File.Close()
		}
		// 确保日志目录存在，如果不存在则创建
		logDir = e.LogRootDir + "/" + currentDay
		if err := os.MkdirAll(e.LogRootDir+"/"+currentDay, os.ModePerm); err != nil {
			return nil, err
		}
		// 创建新的文件夹
		filePath := fmt.Sprintf("%s/%s.log", logDir, currentDay)
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		e.File = file
		e.CurrentDay = currentDay
	}
	e.File.Write(buf.Bytes())

	// err日志输出
	switch entry.Level {
	case zap.ErrorLevel:
		if e.ErrFile == nil {
			file, err := os.OpenFile(logDir+"/"+"err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			if err != nil {
				return nil, err
			}
			e.ErrFile = file
		}
		e.ErrFile.Write(buf.Bytes())
	}

	return buf, nil
}

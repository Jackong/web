/**
 * User: Jackong
 * Date: 13-7-18
 * Time: 下午10:00
 */
package log

import (
	"os"
	"fmt"
	"log"
	"time"
)

type Level string

const (
	DEBUG Level = "Debug"
	INFO Level = "Info"
	WARNING Level = "Warning"
	ERROR Level = "Error"
)

var (
	dir string
	logger *log.Logger
	logFile *os.File
)

func Init(logDir string) {
	dir = logDir
	if err := os.MkdirAll(dir, os.ModeDir); err != nil {
		fmt.Println("warning:", err)
	}
}

func getLog() *log.Logger {
	fileName := dir + "/" + time.Now().Format("2006-01-02") + ".log"
	if logFile != nil {
		if logFile.Name() == fileName {
			return logger
		}
		logFile.Sync()
		logFile.Close()
		logFile = nil
	}

	var err error = nil
	logFile, err = os.OpenFile(fileName, os.O_RDWR | os.O_CREATE, 0)
	if err != nil {
		fmt.Println("warning:", err)
	}
	logger = log.New(logFile, "", log.Ldate | log.Ltime | log.Llongfile)
	return logger
}

func Debug(v ... interface {}) {
	Print(DEBUG, v)
}

func Info(v ... interface {}) {
	Print(INFO, v)
}

func Warning(v ... interface {}) {
	Print(WARNING, v)
}

func Error(v ... interface {}) {
	Print(ERROR, v)
}

func Print(level Level, v ... interface {}) {
	getLog().Print("[",  level, "]|", v)
}

func Close() {
	logFile.Close()
}

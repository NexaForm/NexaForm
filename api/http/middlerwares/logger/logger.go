package logger

import (
	"NexaForm/config"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	ID        string
	Timestamp time.Time
	Level     Level
	Service   string
	Endpoint  string
	UserID    string
	Message   string
	Context   []interface{}
	config    config.Log
}

type Level string

const (
	_DEBUG   Level = "Debug"
	_INFO    Level = "Info"
	_WARNING Level = "Warning"
	_ERROR   Level = "Error"
	_FATAL   Level = "Fatal"
)

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

func New(service string) *Logger {
	l := Logger{
		Service: service,
	}
	return &l
}
func messageMaker(l *Logger) string {
	color := map[Level]string{
		_DEBUG:   blue,
		_INFO:    green,
		_WARNING: yellow,
		_ERROR:   red,
		_FATAL:   red,
	}
	var logMessage strings.Builder
	logMessage.WriteString(color[l.Level] + "\n======> " + string(l.Level) + " <======\n" + reset)
	currentTime := time.Now().UTC()
	isoTimeStr := currentTime.Format(time.RFC3339)
	logMessage.WriteString(isoTimeStr)
	logMessage.WriteString(" " + l.Service)
	if l.UserID != "" {
		logMessage.WriteString(" " + l.UserID)
	}
	logMessage.WriteString(" " + l.Endpoint + ":\n" + color[l.Level] + l.Message + reset + "\n")
	if len(l.Context) > 0 {
		for i, d := range l.Context {
			logMessage.WriteString(fmt.Sprintf("\n%d : %v\n", i, d))
		}
	}
	logMessage.WriteString(magenta + "ID: " + l.ID + reset)
	logMessage.WriteString(color[l.Level] + "\n======> END <======\n" + reset)

	return logMessage.String()
}

func addToFile() {}
func addToDB()   {}

func (l *Logger) SetUser(userID string) *Logger {
	l.UserID = userID
	return l
}
func (l *Logger) Debug(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _DEBUG
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
	fmt.Println(l.config.Address)
}
func (l *Logger) Info(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _INFO
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
}
func (l *Logger) Warning(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _WARNING
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
}
func (l *Logger) Error(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _ERROR
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
}
func (l *Logger) Fatal(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _FATAL
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
}

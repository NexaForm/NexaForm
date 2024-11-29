package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Logger struct {
	ID        string        `json:"id"`
	Timestamp time.Time     `json:"timestamp"`
	Level     Level         `json:"level"`
	Service   string        `json:"service"`
	Endpoint  string        `json:"endpoint"`
	UserID    string        `json:"user_id"`
	Message   string        `json:"message"`
	Context   []interface{} `json:"context"`
	// config    config.Log
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

func appendToFile(logger *Logger) {
	// Create logs directory if it doesn't exist
	os.Mkdir("logs", 0777)
	// Open the log file. This creates the file if it doesn't exist, and opens it in append, create, and read/write mode.
	file, err := os.OpenFile(fmt.Sprintf("./logs/%v_logs.json", logger.Level), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	// Read existing data from the file
	data, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
	}
	// Unmarshal JSON data into a slice of Logger structs
	var jsonData []Logger
	if err := json.Unmarshal(data, &jsonData); err != nil {
		log.Println(err)
	}
	// Append the new log entry to the slice
	jsonData = append(jsonData, *logger)
	// Marshal the updated slice back into JSON
	fileData, err := json.Marshal(jsonData)
	if err != nil {
		log.Println(err)
	}
	// Truncate the file to clear existing content
	file.Truncate(0)
	// Seek to the beginning of the file
	file.Seek(0, 0)
	// Write the updated JSON data to the file
	file.Write(fileData)

}
func appendToDB() {
	//TODO
}

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
	appendToFile(l)
}
func (l *Logger) Info(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _INFO
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Warning(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _WARNING
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Error(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _ERROR
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Fatal(endpoint, message string, data ...interface{}) {
	l.ID = uuid.New().String()
	l.Level = _FATAL
	l.Endpoint = endpoint
	l.Message = message
	l.Context = data
	fmt.Println(messageMaker(l))
	appendToFile(l)
}

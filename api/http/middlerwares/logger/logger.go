package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Logger struct {
	ID        string        `json:"id"`
	Timestamp time.Time     `json:"timestamp"`
	Ip        string        `json:"ip"`
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
	currentTime := time.Now().UTC().Format(time.RFC3339)
	logMessage.WriteString(currentTime)
	logMessage.WriteString("  " + l.Service)
	if l.UserID != "" {
		logMessage.WriteString("  " + l.UserID)
	}
	if l.Ip != "" {
		logMessage.WriteString("  " + l.Ip)
	}
	logMessage.WriteString("  " + l.Endpoint + ":\n" + color[l.Level] + l.Message + reset + "\n\n")
	if len(l.Context) > 0 {
		for i, c := range l.Context {
			if i == 0 {
				logMessage.WriteString(fmt.Sprintf("\t%s: %v\n", "Request", c))
			} else if i == 1 {
				logMessage.WriteString(fmt.Sprintf("\t%s: %v\n", "Response", c))
			}
		}
	}
	logMessage.WriteString("\n\n" + magenta + "ID: " + l.ID + reset)
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

func (l *Logger) SetContext(ctx *fiber.Ctx) *Logger {
	if id := ctx.UserContext().Value("id"); id != nil {
		l.UserID = id.(string)
	}
	l.Endpoint = ctx.Path()
	l.Ip = ctx.IP()
	l.Context = append(l.Context, ctx.Context().Request.Body())
	l.Context = append(l.Context, ctx.Context().Response.Body())
	return l
}
func (l *Logger) Debug(message string) {
	l.ID = uuid.New().String()
	l.Level = _DEBUG
	l.Message = message
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Info(message string) {
	l.ID = uuid.New().String()
	l.Level = _INFO
	l.Message = message
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Warning(message string) {
	l.ID = uuid.New().String()
	l.Level = _WARNING
	l.Message = message
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Error(message string) {
	l.ID = uuid.New().String()
	l.Level = _ERROR
	l.Message = message
	fmt.Println(messageMaker(l))
	appendToFile(l)
}
func (l *Logger) Fatal(message string) {
	l.ID = uuid.New().String()
	l.Level = _FATAL
	l.Message = message
	fmt.Println(messageMaker(l))
	appendToFile(l)
}

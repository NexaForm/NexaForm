package service

import (
	"NexaForm/config"
	"context"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	ServiceAuth     = "auth"
	ServiceDatabase = "database"
	ServiceAPI      = "api"
	ServiceLogger   = "logger"
	ServiceUser     = "user"
)

type LoggerService struct {
	loggers map[string]*zap.Logger // Map for different service loggers
}

// NewLoggerService initializes loggers for multiple services.
func NewLoggerService(logConfigs []config.LoggerConfig, lokiURL string) (*LoggerService, error) {
	loggers := make(map[string]*zap.Logger)

	for _, logConfig := range logConfigs {
		fileWriter := zapcore.AddSync(&lumberjack.Logger{
			Filename:   logConfig.LogFilePath,
			MaxSize:    logConfig.MaxSize,
			MaxBackups: logConfig.MaxBackups,
			MaxAge:     logConfig.MaxAge,
			Compress:   logConfig.Compress,
		})

		atomicLevel := zap.NewAtomicLevel()
		if err := atomicLevel.UnmarshalText([]byte(logConfig.Level)); err != nil {
			return nil, err
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			fileWriter,
			atomicLevel,
		)

		loggers[logConfig.Name] = zap.New(core)
	}

	return &LoggerService{loggers: loggers}, nil
}

// LogInfo for specific service logger
func (l *LoggerService) LogInfo(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Info(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// LogError for specific service logger
func (l *LoggerService) LogError(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Error(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// LogDebug logs debug messages for a specific service
func (l *LoggerService) LogDebug(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Debug(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// LogWarn logs warning messages for a specific service
func (l *LoggerService) LogWarn(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Warn(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// LogFatal logs fatal messages for a specific service and exits the program
func (l *LoggerService) LogFatal(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Fatal(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// LogPanic logs panic messages for a specific service and panics
func (l *LoggerService) LogPanic(ctx context.Context, serviceName, message string, fields ...zap.Field) {
	if logger, exists := l.loggers[serviceName]; exists {
		// Add `service` field for Promtail compatibility
		fields = append(fields, zap.String("service", serviceName))
		logger.Panic(message, append(l.getContextFields(ctx), fields...)...)
	}
}

// getContextFields extracts fields from context (for Fiber context)
func (l *LoggerService) getContextFields(ctx context.Context) []zap.Field {
	var fields []zap.Field

	// Check if the context is a Fiber context
	if fiberCtx, ok := ctx.Value("fiber_ctx").(*fiber.Ctx); ok {
		fields = append(fields,
			zap.String("method", fiberCtx.Method()),
			zap.String("url", fiberCtx.OriginalURL()),
			zap.String("ip", fiberCtx.IP()),
			zap.String("user_agent", fiberCtx.Get("User-Agent")),
		)

		// Add user information if available in Fiber locals
		if userID := fiberCtx.Locals("user_id"); userID != nil {
			fields = append(fields, zap.String("user_id", userID.(string)))
		}
		if email := fiberCtx.Locals("email"); email != nil {
			fields = append(fields, zap.String("email", email.(string)))
		}
	}

	return fields
}

// Sync flushes any buffered log entries
func (l *LoggerService) Sync() {
	// This will flush all log entries to the underlying writers (log files)
	for _, logger := range l.loggers {
		_ = logger.Sync()
	}
}

// GetLoggerFields returns common logging fields
func GetLoggerFields(serviceName string) []zap.Field {
	return []zap.Field{
		zap.String("service", serviceName), // Adding service name as a field
	}
}

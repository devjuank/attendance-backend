package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger inicializa el logger de Zap según el entorno
func InitLogger(env string, logLevel string) error {
	var config zap.Config

	if env == "production" {
		// Configuración para producción (JSON estructurado)
		config = zap.NewProductionConfig()
	} else {
		// Configuración para desarrollo (más legible)
		config = zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	// Configurar nivel de log
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		level = zapcore.InfoLevel
	}
	config.Level = zap.NewAtomicLevelAt(level)

	// Crear logger
	logger, err := config.Build(
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
	if err != nil {
		return err
	}

	Log = logger
	return nil
}

// Sync sincroniza el logger (llamar antes de cerrar la aplicación)
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Helper functions para logging más conveniente

func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// Helper para logging de requests HTTP
func LogRequest(method, path string, statusCode int, duration int64, fields ...zap.Field) {
	baseFields := []zap.Field{
		zap.String("method", method),
		zap.String("path", path),
		zap.Int("status", statusCode),
		zap.Int64("duration_ms", duration),
	}
	baseFields = append(baseFields, fields...)

	if statusCode >= 500 {
		Error("HTTP Request", baseFields...)
	} else if statusCode >= 400 {
		Warn("HTTP Request", baseFields...)
	} else {
		Info("HTTP Request", baseFields...)
	}
}

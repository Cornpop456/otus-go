package main

import (
	"flag"
	"log"
	"os"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/config"
	"github.com/Cornpop456/otus-go/calendar-app/internal/pkg/memstorage"
	"github.com/Cornpop456/otus-go/calendar-app/internal/server"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	configPath string
	logger     *zap.SugaredLogger
)

func init() {
	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()
}

func setupLogger(config *config.Config, logOut *os.File) {
	var logLevel zapcore.Level

	switch config.LogLevel {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.WarnLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	writerSyncer := zapcore.AddSync(logOut)
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, writerSyncer, logLevel)
	logger = zap.New(core, zap.AddCaller()).Sugar()
}

func main() {
	configStruct := &config.Config{}

	if err := configStruct.FromFile(configPath); err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(configStruct.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("opening file error: %v", err)
	}

	defer logFile.Close()

	setupLogger(configStruct, logFile)

	service := server.New(calendar.New(memstorage.NewEventsLocalStorage()), logger)

	if err := service.StartServer(configStruct); err != nil {
		logger.Fatalf("server error: %s", err)
	}
}

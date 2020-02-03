package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar"
	"github.com/Cornpop456/otus-go/calendar-app/internal/config"
	"github.com/Cornpop456/otus-go/calendar-app/internal/pkg/memstorage"

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

	calendar := calendar.New(memstorage.NewEventsLocalStorage())

	t, err := time.Parse(time.RFC822, "26 Jan 20 19:00 MSK")
	if err != nil {
		logger.Fatalf("Parsing err %v", err)
	}

	t2, err := time.Parse(time.RFC822, "25 Feb 21 11:30 MSK")
	if err != nil {
		logger.Fatalf("Parsing err %v", err)
	}

	if _, err := calendar.AddEvent("Do cooking", "cook dinner", t2); err != nil {
		logger.Fatal(err)
	}

	if _, err := calendar.AddEvent("Go to party", "hang out!", t); err != nil {
		logger.Fatal(err)
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("New request from (%s): [METHOD: %s | URL: %s]", r.RemoteAddr, r.Method, r.URL)
		for _, v := range calendar.GetEvents() {
			fmt.Fprintf(w, "Event: %s\nDescription: %s\nDate: %s\n\n", v.Name, v.Description, v.EventDate.String())
		}
	})

	err = http.ListenAndServe(configStruct.HTTPListen, nil)

	logger.Info(err)
}

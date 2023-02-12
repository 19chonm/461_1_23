package logger

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func test() {
	os.Setenv("LOG_FILE", "C:/Users/mmcho/OneDrive/Documents/dummy/461_1_23/")
	os.Setenv("LOG_LEVEL", "1")
}

func init() {
	test()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	path, ok := os.LookupEnv("LOG_FILE")
	if !ok {
		fmt.Println("Location for log file not found")
		os.Exit(1)
	}

	file, err := os.Create(path + "log.json")
	if err != nil {
		fmt.Println("Failed to create log file")
		os.Exit(1)
	}

	log.SetOutput(file)
	logLvl := os.Getenv(("LOG_LEVEL"))

	if logLvl == "1" {
		log.SetLevel(log.InfoLevel)
	} else if logLvl == "2" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.TraceLevel)
	}
}

func DebugMsg(msg ...string) {
	if log.GetLevel() == log.DebugLevel {
		log.Debug(msg)
	}
}

func InfoMsg(msg ...string) {
	if log.GetLevel() == log.InfoLevel {
		log.Info(msg)
	}
}

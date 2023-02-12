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
	path, ok := os.LookupEnv("LOG_FILE")
	if !ok {
		log.SetLevel(log.TraceLevel)
	}

	file, err := os.Create(path + "log.json")
	if err != nil {
		fmt.Println("Failed to create logfile")
	}
	log.SetOutput(file)

	logLvl := os.Getenv(("LOG_LEVEL"))

	if logLvl == "1" {
		log.SetLevel(log.InfoLevel)
	} else if logLvl == "2" {
		log.SetLevel(log.DebugLevel)
	}
}

func DebugMsg(msg ...string) {
	log.Debug(msg)
}

func InfoMsg(msg ...string) {
	log.Info(msg)
}

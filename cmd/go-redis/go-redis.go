package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Selahattinn/go-redis/pkg/version"
	"github.com/sirupsen/logrus"
)

var (
	versionFlag = flag.Bool("version", false, "Show version information.")
	debugFlag   = flag.Bool("debug", false, "Show debug information.")
	addrFlag    = flag.String("addr", ":8080", "Listen address")
	logFileFlag = flag.String("log", "go-redis", "Path to the log file.")
)

func main() {
	flag.Parse()

	if *versionFlag {
		fmt.Fprintln(os.Stdout, version.Print("go-redis"))
		os.Exit(0)
	}

	// Log settings
	if *debugFlag {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetReportCaller(false)
		logrus.SetLevel(logrus.InfoLevel)
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile(*logFileFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Could not open log file")
	}

	// logrus log file setted
	logrus.SetOutput(logFile)
	fmt.Println("Hello world")
}

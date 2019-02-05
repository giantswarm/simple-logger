package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

var logFile = "/etc/slog/creator.log"

func init() {

	if l, present := os.LookupEnv("SIDECAR_LOGFILE_PATH"); present {
		logFile = l
	}

	fmt.Println("Creating", logFile)

	err := os.MkdirAll(filepath.Dir(logFile), os.ModePerm)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(f)
	// log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)
}

func main() {
	fmt.Println("Simple log creator")

	i := 0
	for {
		logline := "The log's number increased tremendously!"
		fmt.Println("Writing '", logline, "' to ", logFile)
		log.WithFields(log.Fields{
			"omg":    true,
			"number": i,
		}).Warn(logline)
		time.Sleep(time.Second * 10)
		i++
	}

}

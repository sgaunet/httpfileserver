package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// initTrace initialize log instance with the level in parameter
func initTrace(debugLevel string) {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	// log.SetFormatter(&log.TextFormatter{
	// 	DisableColors: true,
	// 	FullTimestamp: true,
	// })

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	switch debugLevel {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.DebugLevel)
	}
}

func main() {
	var dirToParse string
	var port int
	initTrace("info")

	flag.StringVar(&dirToParse, "d", "", "Directory to parse")
	flag.IntVar(&port, "p", 8081, "Port of the webserver")
	flag.Parse()

	if dirToParse == "" {
		log.Errorln("specify a directory to expose")
		os.Exit(1)
	}
	if port < 1024 {
		log.Errorln("Port cannot be under 1024")
		os.Exit(1)
	}

	addr := fmt.Sprintf(":%d", port)
	fs := http.FileServer(http.Dir(dirToParse))
	log.Fatal(http.ListenAndServe(addr, fs))
}

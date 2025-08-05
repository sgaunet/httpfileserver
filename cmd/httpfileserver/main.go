// Package main implements a simple HTTP file server with optional basic authentication.
package main

import (
	"flag"
	"os"

	"github.com/sgaunet/httpfileserver/internal/config"
	"github.com/sgaunet/httpfileserver/internal/logger"
	"github.com/sgaunet/httpfileserver/internal/server"
)

func main() {
	var dirToParse string
	var port int
	
	log := logger.New("info")

	flag.StringVar(&dirToParse, "d", "", "Directory to parse")
	flag.IntVar(&port, "p", config.DefaultPort, "Port of the webserver")
	flag.Parse()

	cfg, err := config.NewConfig(dirToParse, port)
	if err != nil {
		log.Errorln(err)
		os.Exit(1)
	}

	srv := server.NewServer(cfg, log)
	log.Fatal(srv.ListenAndServe())
}
package main

import (
	"crypto/sha256"
	"crypto/subtle"
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

type App struct {
	DirToParse string
	User       string
	Password   string
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

	app := App{
		DirToParse: dirToParse,
		User:       os.Getenv("HTTP_USER"),
		Password:   os.Getenv("HTTP_PASSWORD"),
	}

	if app.User != "" && app.Password != "" {
		log.Infoln("Launch webserver with basic auth")
		http.Handle("/", app.basicAuth(app.exposeDir()))
		log.Fatal(http.ListenAndServe(addr, nil))
	} else {
		log.Infoln("Launch webserver without auth")
		fs := http.FileServer(http.Dir(dirToParse))
		log.Fatal(http.ListenAndServe(addr, fs))
	}
}

func (a *App) exposeDir() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(a.DirToParse)).ServeHTTP(w, r)
	})
}
func (a *App) basicAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			usernameHash := sha256.Sum256([]byte(username))
			passwordHash := sha256.Sum256([]byte(password))
			expectedUsernameHash := sha256.Sum256([]byte(a.User))
			expectedPasswordHash := sha256.Sum256([]byte(a.Password))

			usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsernameHash[:]) == 1)
			passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasswordHash[:]) == 1)

			if usernameMatch && passwordMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

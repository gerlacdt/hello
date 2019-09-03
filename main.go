package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/braintree/manners"
	"github.com/gerlacdt/hello/handlers"

	"os/signal"
	"syscall"
)

func logEnvionmentVariables() {
	for _, e := range os.Environ() {
		log.Printf("%s\n", e)
	}
}

func main() {
	log.Printf("Starting hello-app, version: %s\n", handlers.Version)

	logEnvionmentVariables()

	httpPort := os.Getenv("NOMAD_PORT_http")
	if httpPort == "" {
		log.Fatal("NOMAD_PORT_http must be set and not-empty")
	}

	log.Printf("HTTP service port listening on %s", httpPort)

	// create web-router
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(handlers.HelloHandler))
	mux.Handle("/version", http.HandlerFunc(handlers.VersionHandler))
	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"msg\": \"NOT_FOUND\"}")
	})

	// start server
	httpServer := manners.NewServer()
	httpServer.Addr = ":" + httpPort
	httpServer.Handler = mux

	// handle graceful shutdown
	errChan := make(chan error, 10)

	go func() {
		errChan <- httpServer.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Caputured %v. Exiting...", s))
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}

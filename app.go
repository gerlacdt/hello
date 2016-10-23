package main

import (
	"fmt"
	"github.com/braintree/manners"
	"github.com/gerlacdt/hello/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func logEnvionmentVariables() {
	for _, e := range os.Environ() {
		log.Printf("%s\n", e)
	}
}

func main() {
	log.Println("Starting hello-app...")

	logEnvionmentVariables()

	httpAddr := os.Getenv("NOMAD_ADDR_http")
	if httpAddr == "" {
		log.Fatal("NOMAD_ADDR_http must be set and not-empty")
	}

	log.Printf("HTTP service listening on %s", httpAddr)

	// create web-router
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(handlers.HelloHandler))
	mux.Handle("/hello/version", http.HandlerFunc(handlers.VersionHandler))
	mux.Handle("/hello/health", http.HandlerFunc(handlers.HealthHandler))
	mux.Handle("/version", http.HandlerFunc(handlers.VersionHandler))
	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))

	// start server
	httpServer := manners.NewServer()
	httpServer.Addr = httpAddr
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

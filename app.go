package main

import (
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

	// handle SIGTERM
	c := make(chan os.Signal)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Fatalf("Cleaning up before killing process...")
	}()

	// create web-router
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(handlers.HelloHandler))
	mux.Handle("/version", http.HandlerFunc(handlers.VersionHandler))
	mux.Handle("/health", http.HandlerFunc(handlers.HealthHandler))

	// start server
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}

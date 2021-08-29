package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pApi/handlers"
	"syscall"
	"time"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	ph := handlers.NewProduct(l)

	sm := http.NewServeMux()

	sm.Handle("/", hh)
	sm.Handle("/p", ph)

	s := &http.Server{
		Addr:         ":7070",
		Handler:      sm,
		IdleTimeout:  time.Second * 120,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown type :", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}

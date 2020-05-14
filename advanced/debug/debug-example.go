package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inside handler")
	fmt.Println()
}

func main() {
	fmt.Println("starting server")
	srv := http.Server{Addr: "8080",
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	http.HandleFunc("/", handler)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		srv.ListenAndServe()
	}()
	<-done
	err := srv.Shutdown(context.Background())
	if err != nil{
		log.Fatal(err)
	}
}

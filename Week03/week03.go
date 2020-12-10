package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{}))

	go func() {
		done <- ServerApp(stop)
	}()

	go func() {
		done <- ServerApp1(stop)
	}()

	for i := 0; i < cap(done); i++ {
		<-done
		close(stop)
	}
}

func ServerApp(stop <-chan struct{}) error {
	var srv = http.Server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello, serverapp!")
	}

	go func() {
		<-stop
		srv.Shutdown(context.Background())
	}

	http.HandleFunc("/", helloHandler)
	return http.ListenAndServe(":8080", nil)
}

func ServerApp1(stop <-chan struct{}) error {
	var srv = http.Server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, "Hello, deng!")
	}

	go func() {
		<-stop
		srv.Shutdown(context.Background())
	}

	//ceshi
	http.HandleFunc("/hello", helloHandler)
	return http.ListenAndServe(":9090", nil)
}

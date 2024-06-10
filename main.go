package main

import (
	"fmt"
	"handlemod/handlers"
	"net/http"
)

func main() {
	sm := http.NewServeMux()
	handleFunc := &handlers.Products{}
	sm.Handle("/", handleFunc)
	server := http.Server{
		Addr:    ":8080",
		Handler: sm,
	}
	//test
	server.ListenAndServe()
	fmt.Println("server start successa")
}

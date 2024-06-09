package main

import (
	"github.com/gorilla/mux"
	"handlemod/handlers"
	"net/http"
)

func main() {
	//sm := http.NewServeMux()
	//handleFunc := &handlers.Products{}
	//sm.Handle("/", handleFunc)
	//server := http.Server{
	//	Addr:    ":8080",
	//	Handler: sm,
	//}
	//server.ListenAndServe()
	r := mux.NewRouter()
	hnd := &handlers.Products{}
	r.Handle("/", hnd)
	http.ListenAndServe(":8080", r)
}

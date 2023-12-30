package main

import (
	"net/http"
)


func main() {
	mux := http.NewServeMux()
 	file := http.FileServer(http.Dir("./"))
 	mux.Handle("/", file)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe() 
	   
}
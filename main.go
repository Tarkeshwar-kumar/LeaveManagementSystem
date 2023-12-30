package main

import (
	"net/http"
)
func showHomePage(w )

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
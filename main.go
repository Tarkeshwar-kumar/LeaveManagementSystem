package main

import (
	"html/template"
	"net/http"
	
)

func main() {
	mux := http.NewServeMux()
 	file := http.FileServer(http.Dir("./ui/"))
 	mux.Handle("/", file)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/my-leaves", leaveRecords)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe() 
	   
}
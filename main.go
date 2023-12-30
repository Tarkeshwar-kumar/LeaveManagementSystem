package main

import (
	"html/template"
	"net/http"
	"github.com/Tarkeshwar-kumar/LeaveManagementSystem/Handler/handler"
)

func main() {
	mux := http.NewServeMux()
 	file := http.FileServer(http.Dir("./ui/"))
 	mux.Handle("/", file)
	mux.HandleFunc("/login", handlers.Login)
	mux.HandleFunc("/my-leaves", handlers.LeaveRecords)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe() 
	   
}
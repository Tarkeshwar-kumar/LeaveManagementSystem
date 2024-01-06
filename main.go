package main

import (
	"net/http"
	"github.com/Tarkeshwar-kumar/LeaveManagementSystem/handler"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", handler.Login)
	mux.HandleFunc("/apply-for-leave", handler.ApplyForLeave)
	mux.HandleFunc("/my-leaves", handler.LeaveRecords)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe() 
	   
}
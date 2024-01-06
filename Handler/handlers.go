package handler

import (
	"fmt"
	"net/http"
	"html/template"
)

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/index.html")
	tmpl.Execute(w, nil)
	http.Redirect(w, r, "/apply-for-leave", http.StatusSeeOther)
}

func ApplyForLeave(w http.ResponseWriter, r *http.Request){
	fmt.Println("login header", r)
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_application.html")
	tmpl.Execute(w, nil)
}

func LeaveRecords(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	fmt.Println("Leave handler", r)
	if r.FormValue("loginPassword") != "soujanya" {
		return
	}
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_records.html")
	tmpl.Execute(w, nil)
}
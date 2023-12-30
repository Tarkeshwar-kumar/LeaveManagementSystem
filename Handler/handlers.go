package handler

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_application.html")
	tmpl.Execute(w, nil)
}

func LeaveRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_records.html")
	tmpl.Execute(w, nil)
}
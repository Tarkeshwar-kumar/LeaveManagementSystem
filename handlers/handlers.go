package handlers

func login(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_application.html")
	tmpl.Execute(w, nil)
}

func leaveRecords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	tmpl, _ := template.ParseFiles("./ui/leave_records.html")
	tmpl.Execute(w, nil)
}
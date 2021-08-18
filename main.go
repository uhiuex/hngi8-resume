package main

import (
	"html/template"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Default port if not specified
	}
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	name := Name{"Adeyemi Mewayewon", "Adeyemi Mewayewon"}
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, name)
}

type Name struct {
	FName, LName string
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		t, _ := template.ParseFiles("processor.html")
		t.Execute(w, nil)
	}

	fname := r.FormValue("user")
	lname := r.FormValue("email")
	sname := r.FormValue("subject")
	mname := r.FormValue("message")

	d := struct {
		User    string
		Email   string
		Subject string
		Message string
	}{
		User:    fname,
		Email:   lname,
		Subject: sname,
		Message: mname,
	}
	tpl.ExecuteTemplate(w, "processor.html", d)
}

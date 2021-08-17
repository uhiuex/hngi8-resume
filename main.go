package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloworld)
	http.ListenAndServe(":80", nil)
}
func helloworld(w http.ResponseWriter, r *http.Request) {
	name := Name{"Adeyemi Mewayewon"}
	template, _ := template.ParseFiles("index.html")
	template.Execute(w, name)
}

type Name struct {
	FName string
}

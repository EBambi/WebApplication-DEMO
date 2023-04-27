package main

import (
	"html/template"
	"log"
	"net/http"
)

type Greeting struct {
	Name string
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/greet", handleGreet)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	greeting := Greeting{Name: name}
	tmpl := template.Must(template.ParseFiles("greet.html"))
	if err := tmpl.Execute(w, greeting); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

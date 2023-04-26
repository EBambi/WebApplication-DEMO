package main

import (
	"fmt"
	"log"
	"net/http"
)

func webPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi team!")
}

func main() {
	http.HandleFunc("/", webPage)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

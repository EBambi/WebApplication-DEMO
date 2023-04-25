package main

import (
	"fmt"
	"net/http"
)

func webPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test")
}

func main() {
	http.HandleFunc("/", webPage)
	http.ListenAndServe("", nil)
}

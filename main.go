package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type TemplateData struct {
	InputNumbers  string
	SortedNumbers []int
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		input := strings.TrimSpace(r.Form.Get("numbers"))
		numbers := parseNumbers(input)
		sortedNumbers := orderNumbers(numbers)

		data := TemplateData{
			InputNumbers:  input,
			SortedNumbers: sortedNumbers,
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, data)
	}
}

func parseNumbers(input string) []int {
	numStrs := strings.Split(input, ",")
	numbers := make([]int, 0, len(numStrs))
	for _, numStr := range numStrs {
		num, err := strconv.Atoi(strings.TrimSpace(numStr))
		if err != nil {
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func orderNumbers(numbers []int) []int {
	sortedNumbers := make([]int, len(numbers))
	copy(sortedNumbers, numbers)
	sort.Ints(sortedNumbers)
	return sortedNumbers
}

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

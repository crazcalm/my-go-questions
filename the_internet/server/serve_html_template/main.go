package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	helloWorld = `Hello World!
`
)

func main() {
	tmpl, err := template.New("Hello").Parse(helloWorld)
	if err != nil {
		log.Fatalf(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8081", nil)
}

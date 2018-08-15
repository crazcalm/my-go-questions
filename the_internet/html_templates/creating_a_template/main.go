package main

import (
	"html/template"
	"log"
	"os"
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

	tmpl.Execute(os.Stdout, nil)
}

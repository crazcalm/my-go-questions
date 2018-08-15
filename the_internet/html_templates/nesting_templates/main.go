package main

import (
	"html/template"
	"log"
	"os"
)

const (
	templateA = `"I am Template A"`
	templateB = `"I am template B"`
	class     = `Welcome Class!

I would like you to meet template A --> {{template "A"}}
and template B --> {{template "B"}}.
`
)

func main() {
	tmpl, err := template.New("class").Parse(class)
	if err != nil {
		log.Fatalf(err.Error())
	}
	_, err = tmpl.New("A").Parse(templateA)
	if err != nil {
		log.Fatalf(err.Error())
	}

	_, err = tmpl.New("B").Parse(templateB)
	if err != nil {
		log.Fatalf(err.Error())
	}

	tmpl.Execute(os.Stdout, nil)

}

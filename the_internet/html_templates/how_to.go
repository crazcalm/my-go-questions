package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

const (
	helloWorld = `{{define "Hello"}} Hello World!{{end}}`
)

func main() {
	tmpl1, err := template.New("Hello").Parse(helloWorld)
	if err != nil {
		log.Fatalf(err.Error())
	}

	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := tmpl1.New("test").Parse(`{{template "Hello"}}\n{{.Count}} items are made of {{.Material}}\n`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("Templates (tmpl1) %v", tmpl.Templates())
}

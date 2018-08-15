package main

import (
	"html/template"
	"log"
	"os"
)

const (
	introductions = `My name is {{.Name}}!
I am {{.Age}} years old.

{{if .Friends}} My friends include {{range .Friends}}{{.}}, {{end}}and, of course, you all! {{else}}I have no friends...{{end}}

Nice to meet you all!

`
)

type person struct {
	Name    string
	Age     int
	Friends []string
}

func main() {
	marcus := person{"Marcus Willock", 30, []string{"Me", "Myself", "and I"}}
	theRealMarcus := person{"Marcus Willock", 30, []string{}}
	tmpl, err := template.New("intro").Parse(introductions)
	if err != nil {
		log.Fatalf(err.Error())
	}
	tmpl.Execute(os.Stdout, marcus)
	tmpl.Execute(os.Stdout, theRealMarcus)

}

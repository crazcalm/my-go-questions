package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	page = `<!DOCTYPE html>
<html lang="en">
  <head>
  	<link rel="stylesheet" href="/static/style.css">
    <title>Sample style page</title>
  </head>
  <body>
    <h1>Sample style page</h1>
    <p>This page is just a demo</p>
  </body>
</html>
`
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	})
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
}

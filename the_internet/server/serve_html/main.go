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
    <title>Sample style page</title>
    <style>
      body {background: navy; color: yellow;}
    </style>
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
	log.Fatal(http.ListenAndServe(":8081", nil))
}

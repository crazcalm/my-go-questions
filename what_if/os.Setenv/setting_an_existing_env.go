/*
Question:

If you use os.Setenv to override an existing environment variable,
will the new value persists post the process (the one that run os.Setenv)
ending?
*/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	browser := os.Getenv("BROWSER")

	fmt.Printf("Browser's path before setting it: %s\n", browser)

	err := os.Setenv("BROWSER", "fake_path")
	if err != nil {
		log.Fatal(err)
	}

	browser = os.Getenv("BROWSER")

	fmt.Printf("Browser's path post setting it: %s\n", browser)
}

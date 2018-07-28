package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type myFile struct {
	http.File
}

func (f myFile) Readdir(n int) ([]os.FileInfo, error) {
	files, err := f.File.Readdir(n)
	return files, err
}

type myFileSystem struct {
	http.FileSystem
}

func isDotFile(name string) (result bool) {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			result = true
			return
		}
	}
	return
}

func (fs myFileSystem) Open(name string) (http.File, error) {
	fmt.Printf("name: %s\n", name)
	file, err := fs.FileSystem.Open(name)

	if isDotFile(name) {
		return file, os.ErrPermission
	}

	return myFile{file}, err
}

func main() {
	home := os.Getenv("HOME")
	hideDotFileSystem := myFileSystem{http.Dir(home)}
	http.Handle("/", http.FileServer(hideDotFileSystem))
	log.Fatal(http.ListenAndServe(":12346", nil))
}

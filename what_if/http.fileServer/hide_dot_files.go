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

func (f myFile) Readdir(n int) (wantedFiles []os.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files { // Filters out the dot files
		if !strings.HasPrefix(file.Name(), ".") {
			wantedFiles = append(wantedFiles, file)
		}
	}

	return
}

type myFileSystem struct {
	http.FileSystem
}

//isDotFile -- checks to see if name is a dot file
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
	fmt.Printf("myFileSystem Open(%s)\n", name)
	file, err := fs.FileSystem.Open(name)

	if isDotFile(name) { //If dot file return 403 response
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

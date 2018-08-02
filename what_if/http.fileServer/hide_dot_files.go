package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
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

func (fs myFileSystem) Open(name string) (file http.File, err error) {
	fmt.Printf("myFileSystem Open(%s)\n", name)
	if isDotFile(name) { //If dot file return 403 response
		return file, os.ErrPermission
	}
	file, err = fs.FileSystem.Open(name)
	return myFile{file}, err
}

func exitError(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	os.Exit(1)
}

func validPath(path string) {
	_, err := os.Stat(path)
	if err != nil {
		exitError(fmt.Errorf("Could not validate path: %s", err.Error()))
	}
}

func validPort(port string) {
	_, err := strconv.Atoi(port)
	if err != nil {
		exitError(fmt.Errorf("Could not confirm that the passed in port was a number: %s", err.Error()))
	}
}

var filePath = flag.String("directory", os.Getenv("HOME"), "File directory to be served")
var showDotFiles = flag.Bool("dotfiles", false, "Show dot files")
var port = flag.String("port", "12345", "That port that will be used")

func main() {
	flag.Parse()

	//Some validation
	validPath(*filePath)
	validPort(*port)

	home := *filePath
	if *showDotFiles {
		http.Handle("/", http.FileServer(http.Dir(home)))
	} else {
		hideDotFileSystem := myFileSystem{http.Dir(home)}
		http.Handle("/", http.FileServer(hideDotFileSystem))
	}

	//Let the user know where the server is
	fmt.Printf("Now serving on http://localhost:%s\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
}

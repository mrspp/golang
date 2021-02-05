package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/mrspp/golang-api/urls"
)

var (
	fileName    string
	fullURLFile string
)

func main() {

	urls := urls.Urls()
	for _, v := range urls {
		fullURLFile = v
		// fullURLFile = "http://www.golangprograms.com/skin/frontend/base/default/logo.png"

		// Build fileName from fullPath
		buildFileName()

		// Create blank file
		file := createFile()

		// Put content on file
		putFile(file, httpClient())
	}

}

func putFile(file *os.File, client *http.Client) {
	resp, err := client.Get(fullURLFile)

	checkError(err)

	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	checkError(err)

	fmt.Println("Just Downloaded a file %s with size %d", fileName, size)
}

func buildFileName() {
	fileURL, err := url.Parse(fullURLFile)
	checkError(err)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName = segments[len(segments)-1]
}

func httpClient() *http.Client {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	return &client
}

func createFile() *os.File {
	file, err := os.Create(fileName)

	checkError(err)
	return file
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

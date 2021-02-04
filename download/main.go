package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	getImage("https://pkg.go.dev/static/img/go-logo-blue.svg")
}

func getImage(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println(string(resp))
	return string(resp), nil
}

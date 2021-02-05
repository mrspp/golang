package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	URL := "https://picsum.photos/id/1008/5616/3744"
	res, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close() // for garbage collection

	responseBodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return
	}
	val := http.DetectContentType(responseBodyBytes)
	fmt.Println(val)
	tail := strings.Split(val, "/")
	filename := tail[len(tail)-1]
	fmt.Println(filename)

}

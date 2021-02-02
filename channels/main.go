package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func dataSource(ch chan int) {
	for {
		ch <- int(rand.Int63())
		time.Sleep(time.Second)
	}
}

func runLoop(id int32, ch chan string, waitTime int) {
	for {
		process(id, <-ch, waitTime)
	}
}

func process(id int32, s string, waitTime int) {
	time.Sleep(time.Duration(waitTime) * time.Second)
	fmt.Printf("process: %d - %s \n", id, s)
}

func main() {
	fmt.Printf("Start \n")

	myChan := make(chan int)
	myChan1 := make(chan int)
	go dataSource(myChan)
	go dataSource(myChan1)

	dataCh := make(chan int, 1)
	go func() {
		for {
			fmt.Println(<-dataCh)
		}
	}()

	go func() {
		for {
			select {
			case data := <-myChan:
				dataCh <- data
			case data := <-myChan1:
				dataCh <- data
			default:
			}
			// fmt.Println("Hahaha")
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh
	fmt.Println("Hihi")
}

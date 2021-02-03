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

func main() {

	// defer func() {
	// 	r := recover()
	// 	if r != nil {
	// 		fmt.Println("stop hahaha")
	// 	}
	// }()

	fmt.Printf("Start \n")

	myChan := make(chan int)
	myChan1 := make(chan int)

	// defer func() {
	// 	fmt.Println("hihi")
	// }()

	// panic("ssss")

	go dataSource(myChan)
	go dataSource(myChan1)

	dataCh := make(chan int, 1)
	go func() {
		for {
			fmt.Println(<-dataCh)
		}
	}()

	// helper.RunLoop()

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

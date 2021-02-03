package helper

import (
	"fmt"
	"time"
)

func process(id int32, s string, waitTime int) {

	time.Sleep(time.Duration(waitTime) * time.Second)
	fmt.Printf("process: %d - %s \n", id, s)
}

// RunLoop is a function to loop channels
func RunLoop(id int32, ch chan string, waitTime int) {
	for {
		process(id, <-ch, waitTime)
	}
}

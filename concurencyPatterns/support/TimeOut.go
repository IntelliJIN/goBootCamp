package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 1 * time.Second
	c := make(chan int)
	go func(timeout time.Duration) {
		select {
		case m := <-c:
			fmt.Println("Received: ", m)
		case <-time.After(duration):
			fmt.Println("timed out")
		}
	}(duration)
	time.Sleep(5 * time.Second)
}

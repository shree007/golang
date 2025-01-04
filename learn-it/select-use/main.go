package main

import (
	"fmt"
	"time"
)

func main() {
	simpleUseSelect()
	nonBlockingCommunication()
	infiniteLoop()
}

func infiniteLoop() {
	ticker := time.NewTicker(1 * time.Second)
	timeout := time.After(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick...")
		case <-timeout:
			fmt.Println("Time's up!")
			return
		}
	}
}

func nonBlockingCommunication() {
	ch := make(chan string)
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No data received")
	}
}

func simpleUseSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from channel 1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from channel 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Recieved:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recieved:", msg2)
	}
}

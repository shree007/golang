package main


import (
	"fmt"
	"time"
)

func sendData(ch chan int){
	for i:=1; i<=5; i++{
		fmt.Println("sending ", i)
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func recieveData(ch chan int){
	for {
		value, ok := <-ch // Here OK is boolean value which is decision maker
		if !ok {
			fmt.Println("closing channel, existing")
			return
		}
		fmt.Println("Recieved.... ", value)
	}
}

func main(){
	ch := make(chan int)
	go sendData(ch)
	go recieveData(ch)

	time.Sleep(10 * time.Second)
}
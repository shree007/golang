package main


import (
	"fmt"
	"time"
)


func printNumbers(){
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printLetters(){
	for i:= 'a'; i<'e'; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main(){
	fmt.Println("Learning Goroutines")
	go printNumbers()
	printLetters()
	time.Sleep(1 * time.Second)
	print("\n", "Done")
}

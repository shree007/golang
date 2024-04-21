package main

import "fmt"

func example_of_panic(){
	defer fmt.Println(" it should be called before termination")
	fmt.Print("inside into example_of_panic function")
	panic("Something went wrong")

}

func main(){
	fmt.Println("main function started to exeute")
	example_of_panic()
	fmt.Println("This should not be printed if panic concept is true")
}
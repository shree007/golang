package main


import (
	"fmt"
)

func execute_panic(){
	fmt.Println("Inside panic function")
	panic("Something went wrong")
}

func recover_panic(){
	if r:= recover(); r!=nil{
		fmt.Println("Recovered: ", r)
	}
}

func main(){
	fmt.Println("Main start.......")
	defer  recover_panic()
    execute_panic()
    fmt.Println("if it is being printed it means recover handle panic gracefully")
}

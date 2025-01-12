package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an argument")
		return
	}
	fmt.Println("First Argument:", os.Args[1])
	fmt.Println("Second Argument:", os.Args[2])
}

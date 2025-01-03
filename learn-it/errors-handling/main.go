package main

import (
	"errors"
	"fmt"
)

func main() {
	basicErrorHandling()
	customErrorHandling()
}

func customErrorHandling() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	b := 0
	a := 10
	if b == 0 {
		panic(fmt.Sprintf("cannot divide %d by zero", a))
	}
	fmt.Println("division:", a/b)
}

func basicErrorHandling() {
	msg := "division by zero"
	result, err := divide(10, 0, msg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result: ", result)
}

func divide(a, b int, msg string) (int, error) {
	if b == 0 {
		return 0, errors.New(msg)
	}
	return a / b, nil
}

package main

import (
	"errors"
	"fmt"
)

var ErrorNotFound = errors.New("Not Found")

func main() {
	basicErrorHandling()
	customErrorHandling()
	errorsIs()
	errorsAs()
}

type MyCustomError struct {
	Message string
}

func (e *MyCustomError) Error() string {
	return e.Message
}

func triggerError() error {
	return &MyCustomError{Message: "Something went wrong"}
}

func errorsAs() {
	err := triggerError()
	var customErr *MyCustomError

	if errors.As(err, &customErr) {
		fmt.Println("Caught a custom error:", customErr.Message)
	} else {
		fmt.Println("Some other error occurred")
	}
}

func findUser(id int) error {
	if id == 0 {
		return ErrorNotFound
	}
	return nil
}

func errorsIs() {
	err := findUser(0)
	if errors.Is(err, ErrorNotFound) {
		fmt.Println("User Not Found")
	} else {
		fmt.Println("User Found")
	}
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

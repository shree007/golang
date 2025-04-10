package main

import (
	"fmt"
	"log"
	//"log/slog" available: 1.21 version, Learning curve https://betterstack.com/community/guides/logging/logging-in-go/
)

func awsCreateEcrRepository() {
	log.Printf("hello, world")
}

func main() {
	fmt.Println("create docker repository")
	awsCreateEcrRepository()
}

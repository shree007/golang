package main

import (
	"fmt"
	"os"
)

func main() {
	app := os.Getenv("APP_NAME")

	switch app {
	case "app1":
		startServer("App1", "8080")
	case "app2":
		startServer("App2", "8090")
	case "app3":
		startServer("App3", "8081")
	default:
		fmt.Println("No valid APP_NAME provided. Use app1, app2, or app3.")
		os.Exit(1)
	}
}

func startServer(name, port string) {
	fmt.Println("name: ", name)
	fmt.Println("port: ", port)
}

package main

import (
	"fmt"
	"net/http"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello from %s running on port %s \n", name, port)
	})

	fmt.Printf("%s is running on port...\n", name, port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

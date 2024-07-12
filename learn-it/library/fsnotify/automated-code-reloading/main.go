package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func main() {
	sourceDir := "src"
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal()
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("Source file modified:", event.Name)
					cmd := exec.Command("go", "run", sourceDir)
					output, err := cmd.CombinedOutput()
					if err != nil {
						log.Printf("Error running code: %v\nOutput: %s\n", err, output)
					} else {
						log.Printf("Code output: %s\n", output)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Watch error:", err)
			}
		}
	}()
	err = watcher.Add(sourceDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}

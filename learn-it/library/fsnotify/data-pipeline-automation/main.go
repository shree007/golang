package main

import (
    "fmt"
    "log"
    "github.com/fsnotify/fsnotify"
)

func runDataPipeline(event fsnotify.Event) {
    fmt.Printf("Data pipeline triggered by: %s\n", event)
}

func main() {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    directoryToWatch := "data"

    err = watcher.Add(directoryToWatch)
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)
    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write {
                    runDataPipeline(event)
                }

            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    fmt.Println("Watching directory:", directoryToWatch)
    <-done
}
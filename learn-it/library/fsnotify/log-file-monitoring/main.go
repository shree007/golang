package main

import(
	"log"
	"github.com/fsnotify/fsnotify"
)

func main(){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func(){
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok{
					return 
				}
				log.Println("events: ", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file: ", event.Name)
				}
			case err, ok := <-watcher.Events:
				if !ok{
					return 
				}
				log.Println("error:", err)

			}
		}
	}()
	err = watcher.Add("logs")
	if err != nil{
		log.Fatal(err)
	}
	<-make(chan struct{})
}
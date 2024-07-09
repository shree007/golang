package main

import(
	"fmt"
	"io/ioutil"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"log"
)

type Config struct {
	AppName  string `yaml:app_name`
	LogLevel string `yaml:log_level`
	Path     string `yaml:path`
}

var config Config

func loadConfig(){
	data, err := ioutil.ReadFile("etc/service/config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println("Config loaded: %+v\n", config)

}


func main(){
	loadConfig()
	watcher, err := fsnotify.NewWatcher()
	if err != nil{
		log.Fatal(err)
	}
	defer watcher.Close()

	go func(){
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
			   	}
			   	if event.Op&fsnotify.Write == fsnotify.Write {
			   		fmt.Println("Service Config file modified: ", event.Name)
			   		loadConfig()
			   	}
			case err, ok := <-watcher.Errors:
				if !ok{
					return
				}
				fmt.Println("Watcher has error:", err) 	
			}
		}
	}()

	err = watcher.Add("etc/service/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	<-make(chan struct{})
}
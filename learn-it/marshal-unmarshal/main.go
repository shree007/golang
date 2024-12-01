package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

/*
Marshal: convert strcut to JSON/YAML
Unmarshal: convert JSON/YAML to struct
*/

type Person struct {
	Name    string   `yaml:"name"`
	Age     int      `yaml:"age"`
	Hobbies []string `yaml:"hobbies"`
}

func main() {
	person := Person{
		Name:    "shreeprakash",
		Age:     32,
		Hobbies: []string{"books", "movies", "badminton"},
	}
	data, err := yaml.Marshal(person)
	if err != nil {
		log.Fatalf("error while marshalling to yaml %v", err)
	}
	fmt.Println("YAML converted data is:", string(data))

	err = os.WriteFile("person.yaml", data, 0644)
	if err != nil {
		log.Fatalf("Problem while writing into file %v", err)
	}
	fmt.Println("Data has been written into person.yaml file")

}

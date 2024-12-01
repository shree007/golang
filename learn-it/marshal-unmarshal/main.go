package main

import (
	"encoding/json"
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
	marshalYaml()
	unmashalToYaml()

	marshalJson()
	unmarshalJson()
}
func marshalJson() {
	fmt.Println("==========Marshal person struct inJSON==========")
	person := Person{
		Name:    "shreeprakash",
		Age:     32,
		Hobbies: []string{"books", "movies", "badminton"},
	}
	jsonData, err := json.MarshalIndent(person, "", " ")
	if err != nil {
		log.Fatalf("Error while marshal to json")
	}
	fmt.Println("JSON Data:", string(jsonData))
	err = os.WriteFile("person.json", jsonData, 0644)
	if err != nil {
		log.Fatalf("Error while writing struct into person.json")
	}
	fmt.Println("Person struct has been written into person.json")
}

func unmarshalJson() {
	fmt.Println("==========UNMarshal person struct inJSON==========")
	readJsonData, err := os.ReadFile("person.json")
	if err != nil {
		log.Fatalf("Error while reading json file person.json")
	}
	var person Person
	err = json.Unmarshal(readJsonData, &person)
	if err != nil {
		log.Fatalf("Error while unMarshal %v", err)
	}
	fmt.Println("Unmarshaled Struct:\n", person)

}
func marshalYaml() {
	fmt.Println("==========Marshal person struct==========")
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

func unmashalToYaml() {
	fmt.Println("==========Unmarshal person.yaml==========")
	readData, err := os.ReadFile("person.yaml")
	if err != nil {
		log.Fatalf("error while reading person.yaml file %v", err)
	}
	var person Person
	err = yaml.Unmarshal(readData, &person)
	if err != nil {
		log.Errorf("Error while marshalling data to struct %v", err)
	}
	fmt.Println("UnMarshalled struct: %v\n", person)
}

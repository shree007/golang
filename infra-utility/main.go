package main

import (
	"github.com/tidwall/gjson" // more info https://github.com/tidwall/gjson
	"io/ioutil"
	"log"
	"os/exec"
	"os"
	"strings"
	"fmt"
)

func getcurrentdir() string {
  command_to_execute := "pwd"
  cmd := exec.Command(command_to_execute)
  stdout, err := cmd.Output()
  fmt.Println(err)
  return string(stdout)
}


func infrautility(utility_name string){
	// read `versions.json` file
	path := "./"+utility_name+"/versions.json"
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}
    
    alpine_version := gjson.Get(string(content),"utilites.alpine_version")
    log.Printf("%s\n",alpine_version)
    
    command_to_exeecute := "cd " + getcurrentdir() + "/" + utility_name + " && docker build" + " -t " + utility_name +  " ."
    fmt.Println(command_to_exeecute)
    
    cmd := exec.Command(command_to_exeecute)

    stdout, err := cmd.Output()

    if err != nil{
    	fmt.Println(err.Error())
    }

    fmt.Println(string(stdout))

}

func main(){

	utility_name := os.Args[1]

	if strings.Compare(string(utility_name),  "infrautility") == 0 {
 			infrautility(string(utility_name))
	}
}
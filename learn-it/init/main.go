package main


import (
	"fmt"
)

var global_var string

func init(){
	global_var="Hello i am coming from init function"
}

func main(){
	fmt.Println("Main function start")
	fmt.Println(global_var)
}
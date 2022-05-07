package main


import (

	"fmt"
	"github.com/logrusorgru/aurora"
)

func main(){
	greetingMsgEmpty := greeting("")
	fmt.Println(aurora.Green(greetingMsgEmpty))
	greetingMsg := greeting("John")
	fmt.Println(aurora.Yellow(greetingMsg))
}
package main



import "fmt"

func main(){
	defer fmt.Println("I will be Printed later")
	defer fmt.Println("I will be first in defer because of its nature to execute reverser order in multiple defer statement")
	fmt.Println("hola hola hola")
}
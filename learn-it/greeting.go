package main



import (
	"fmt"
)

func greeting(user string) string {
	if(len(user) == 0){
		return "Hello Dude"
	}
	return fmt.Sprintf("Hello %v!", user)
}
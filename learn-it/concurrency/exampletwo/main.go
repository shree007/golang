package main



import (
	"time"
	"fmt"
)

func worker(id int, collector <-chan int, delivers chan <- int){
	for c := range collector {
		fmt.Println("worker", id, "started job", c)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", c)
		delivers <- c * 2
	}
}

func main(){
	number_of_goods := 10
    collector := make(chan int, number_of_goods)
    delivers   := make(chan int, number_of_goods)

    // create number of workers to work concurrently
    for i:=1; i<=3;i++{
    	go worker(i,collector,delivers)
    }

    for c:=1; c<=number_of_goods; c++{
    	collector <- c
    }
    close(collector)

    for d:=1;d<=number_of_goods;d++{
    	deliver:= <-delivers
    	fmt.Println("deliver", deliver)
    }

}
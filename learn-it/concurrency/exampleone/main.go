package main

import (

	"fmt"
	"time"
)


func worker(id int, jobs <-chan int, results chan<- int){
	for j := range jobs{
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main(){
	number_of_jobs := 5
	jobs := make(chan int,number_of_jobs)
	results := make(chan int, number_of_jobs)

	for i:=1; i<=3; i++{
		go worker(i, jobs, results)
	}

	for j:=1; j<=number_of_jobs; j++{
		jobs <- j
	}
	close(jobs)

	for a:=1; a<=number_of_jobs; a++{
		result := <-results
		fmt.Println("Result:", result)

	}
}
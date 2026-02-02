package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ID int, wg *sync.WaitGroup, sharedChannel <-chan int) {
	defer wg.Done()

	for i := range sharedChannel { //listening to channel using range loop
		fmt.Printf("worker Id: %d started the job: %d\n", ID, i)
		time.Sleep(3 * time.Second)//waiting before execution 
		fmt.Printf("worker Id: %d finished the job: %d\n", ID, i)
	}
}

func main() {
	var wg sync.WaitGroup
	workers := 0
	jobs := 0

	fmt.Println("Enter the Total workers: ")
	fmt.Scanln(&workers)

	fmt.Println("Enter the Total Jobs: ")
	fmt.Scanln(&jobs)

	wg.Add(workers)
	
	sharedChannel := make(chan int, jobs)

	for i := 1; i <= workers; i++ {
		go worker(i, &wg, sharedChannel)//calling all workers as goroutine
	}

	for i := 1; i <= jobs; i++ { 
		sharedChannel <- i //writing to the channel
	}

	close(sharedChannel)

	wg.Wait()

	fmt.Println("\nDone Execution!")
}

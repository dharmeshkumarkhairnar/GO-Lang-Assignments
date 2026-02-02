package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ID int, wg *sync.WaitGroup, sharedChannel <-chan int) {
	defer wg.Done()
	for i := range sharedChannel {
		fmt.Printf("worker Id: %d started the job: %d\n", ID, i)
		time.Sleep(3 * time.Second)
		fmt.Printf("worker Id: %d finished the job: %d\n", ID, i)
	}
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Enter the Total workers: ")
	workers := 0
	fmt.Scanln(&workers)

	fmt.Println("Enter the Total Jobs: ")
	jobs := 0
	fmt.Scanln(&jobs)

	wg.Add(workers)
	sharedChannel := make(chan int, jobs)

	for i := 01; i <= workers; i++ {
		go worker(i, &wg, sharedChannel)
	}

	for i := range jobs {
		sharedChannel <- i + 1
	}
	close(sharedChannel)

	wg.Wait()

	fmt.Println("\nDone Execution!")
}

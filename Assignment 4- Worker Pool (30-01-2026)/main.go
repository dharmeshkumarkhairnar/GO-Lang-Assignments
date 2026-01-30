package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func worker(ID int, wg *sync.WaitGroup, sharedChannel <-chan int) {
	defer wg.Done()
	for i := range sharedChannel {
		fmt.Println("")
		fmt.Printf("worker Id: %d doing the job: %d", ID, i)
	}
}

func main() {
	var wg sync.WaitGroup

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the Total workers: ")
	workers, _ := reader.ReadString('\n')
	workers = strings.TrimSpace(workers)
	totalWorkers, _ := strconv.Atoi(workers)

	fmt.Println("Enter the Total Jobs: ")
	jobs, _ := reader.ReadString('\n')
	jobs = strings.TrimSpace(jobs)
	totalJobs, _ := strconv.Atoi(jobs)

	wg.Add(totalWorkers)
	sharedChannel := make(chan int, totalJobs)

	for i := range totalWorkers {
		go worker(i+1, &wg, sharedChannel)
	}

	for i := range totalJobs {
		sharedChannel <- i+1
	}
	close(sharedChannel)

	wg.Wait()

	fmt.Println("")
	fmt.Println("Done Execution!")
	fmt.Println("")
}

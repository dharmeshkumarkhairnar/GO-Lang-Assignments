package main

import (
	"fmt"
	"sync"
)

//Worker which will be run for every element of the list
func squareWorker(element float64, wg *sync.WaitGroup, result chan float64) {
	defer wg.Done()
	val := element * element
	fmt.Printf("Sqaured number is %.2f\n", val)
	result <- val
}

//Funtion which takes the squared values and sum it then pass Total sum of squares through channel.
func aggregateSquares(result chan float64,totalSum chan float64 ) {
	sum:=0.0
	for i := range result {
		sum += i
	}
	totalSum<-sum
}

func main() {
	size := 0
	fmt.Println("Enter the size of list:")
	_, err := fmt.Scanln(&size)

	//input validation
	if err != nil || size < 0 {
		fmt.Println("Enter the Correct input")
		return
	}

	//using the float list so user can enter int as well as float
	arr := make([]float64, size)

	fmt.Println("Enter the numbers:")
	for i := 0; i < size; i++ {
		_, err := fmt.Scanln(&arr[i])
		if err != nil {
			fmt.Println("Enter the correct entry")
			return
		}
	}

	var wg sync.WaitGroup

	//channel for taking sq. values from sqworker and provide to aggregate function.
	result := make(chan float64, size)

	//channel for taking the final result from the aggregate function to main function.
	totalSum := make(chan float64)
	

	go aggregateSquares(result,totalSum)

	wg.Add(size)

	for _, ele := range arr {
		go squareWorker(ele, &wg, result)
	}

	wg.Wait()

	//Closing the range loop as all the worker has finished their job 
	//No values will be sent on this channel
	close(result)

	//recieving the final sum from the aggregate function
	fmt.Printf("Summation of squares is %.2f \n", <-totalSum)
}

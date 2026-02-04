//**********************Using the Unbuffered channel***************** 

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var totalSum int

// func squareWorker(chan1 chan int, wg *sync.WaitGroup, result chan int) {
// 	defer wg.Done()


// 	go aggregateSquares(result,wg)

// 	for i := range chan1 {
// 		val := i * i
// 		fmt.Println("Square of number is", val)
// 		result <- val
// 	}

// 	close(result)

// }

// func aggregateSquares(result chan int, wg *sync.WaitGroup) {
// 	for i := range result {
// 		totalSum += i
// 	}
// 	wg.Done()
// }

// func main() {
// 	size:=0
// 	fmt.Println("Enter the size of list:")
// 	fmt.Scanln(&size)

// 	arr:=make([]int,size)
// 	fmt.Println("Enter the numbers:")
// 	for i:=0;i<size;i++ {
// 		fmt.Scanln(&arr[i])
// 	}

// 	var wg sync.WaitGroup

// 	chan1 := make(chan int)
// 	result := make(chan int)

// 	wg.Add(2)

// 	go squareWorker(chan1, &wg, result)

// 	for _, d := range arr {
// 		chan1 <- d
// 	}
// 	close(chan1)

// 	wg.Wait()

// 	fmt.Println("Summation of squares is",totalSum)

// }



//*****************Using the Buffered channel********************

package main

import (
	"fmt"
	"sync"
)

var totalSum float64

func squareWorker(chan1 chan float64, wg *sync.WaitGroup, result chan float64) {
	defer wg.Done()

	for i := range chan1 {
		val := i * i
		fmt.Printf("Square of number is %.2f\n", val)
		result <- val
	}
	close(result)

}

func aggregateSquares(result chan float64) {
	for i := range result {
		totalSum += i
	}
}

func main() {
	size:=0
	fmt.Println("Enter the size of list:")
	_,err:=fmt.Scanln(&size)

	if err!=nil || size<0 {
		fmt.Println("Enter the Correct input: ", err)
		return
	}

	arr:=make([]float64,size)
	fmt.Println("Enter the numbers:")
	for i:=0;i<size;i++ {
		_,err:=fmt.Scanln(&arr[i])
		if err!=nil {
			fmt.Println("Enter the correct entry")
			return
		}
	}

	var wg sync.WaitGroup

	chan1 := make(chan float64,size)
	result := make(chan float64)

	wg.Add(1)

	go squareWorker(chan1, &wg, result)

	for _, d := range arr {
		chan1 <- d
	}
	close(chan1)

	aggregateSquares(result)

	wg.Wait()

	fmt.Printf("Summation of squares is %.2f \n",totalSum)

}

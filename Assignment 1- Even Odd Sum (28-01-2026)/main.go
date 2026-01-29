//**************************Using the Channels************************

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func Even(list []int, wg *sync.WaitGroup, even chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for val := range list {
// 		if val%2 == 0 {
// 			sum += val
// 		}
// 	}
// 	even <- sum
// }

// func Odd(list []int, wg *sync.WaitGroup, odd chan int) {
// 	defer wg.Done()
// 	sum := 0
// 	for val := range list {
// 		if val%2 == 1 {
// 			sum += val
// 		}
// 	}
// 	odd <- sum
// }

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

// 	var wg sync.WaitGroup

// 	even := make(chan int)
// 	odd := make(chan int)

// 	wg.Add(2)

// 	go Even(arr, &wg, even)
// 	go Odd(arr, &wg, odd)

// 	resEven := <-even
// 	resOdd := <-odd

// 	fmt.Println("Even Sum is ", resEven)
// 	fmt.Println("Odd Sum is ", resOdd)

// 	wg.Wait()

// }
//*********************Using the Pointer variable***************************

package main

import (
	"fmt"
	"sync"
)

func Even(list []int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	localSum := 0
	for val := range list {
		if val%2 == 0 {
			localSum += val
		}
	}
	*sum = localSum
}

func Odd(list []int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	localSum := 0
	for val := range list {
		if val%2 == 1 {
			localSum += val
		}
	}
	*sum = localSum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var wg sync.WaitGroup

	evenSum := 0
	oddSum := 0

	wg.Add(2)

	go Even(arr, &wg, &evenSum)
	go Odd(arr, &wg, &oddSum)

	wg.Wait()

	fmt.Println("Even Sum is ", evenSum)
	fmt.Println("Odd Sum is ", oddSum)

}

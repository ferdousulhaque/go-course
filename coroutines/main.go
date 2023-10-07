package main

import (
	"fmt"
)

func sum(nums []int, ch chan<- int) {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// fmt.Println("Summation: ", sum)
	ch <- sum
}

func main() {
	nums := []int{1, 5, 7, 9, 11}

	// Unbuffered Channel
	ch := make(chan int)

	go sum(nums, ch)
	result := <-ch

	fmt.Println("Summation: ", result)

	// time.Sleep(100 * time.Millisecond)

	//Buffered Channel
	ch2 := make(chan string, 2)

	ch2 <- "Ferdous"
	ch2 <- "haque"

	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
}

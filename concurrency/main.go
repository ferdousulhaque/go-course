package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fill(ch chan<- int, max int) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	for i := 0; i < cap(ch); i++ {
		ch <- r.Intn(max)
	}

	close(ch)
}

func main() {
	numChan := make(chan int, 5)

	go fill(numChan, 100)

	for num := range numChan {
		fmt.Println(num)
	}
}

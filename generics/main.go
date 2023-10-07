package main

import "fmt"

func sum[T int | float64](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(sum(50, 50))
	fmt.Println(sum(1.25, 3.75))
}

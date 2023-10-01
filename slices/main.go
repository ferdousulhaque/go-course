package main

import "fmt"

func main() {
	// declare a slice using make function
	var slice = make([]string, 0)

	slice = append(slice, "One")
	slice = append(slice, "Two")
	slice = append(slice, "Three")
	slice = append(slice, "Four")

	fmt.Println(slice)

	// slice the slice using [low:high] bounds
	fmt.Println(slice[1:3])
	fmt.Println(slice[1:])
	fmt.Println(slice[:1])
	fmt.Println(slice[:])
}

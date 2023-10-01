package main

import "fmt"

// declare a global variable val
var val = "global"

func main() {
	// declare a local variable val and assign integer
	val := 42
	// print the type of local variable and value
	fmt.Printf("Type %T and value %v\n", val, val)
	// print the value of global variable
	printGlobal()
	// update global variable and print
	updateGlobal()
	printGlobal()
	// print the memory location of local variable val
	fmt.Printf("%T, local &val = %v\n", &val, &val)
	// Update the local val using pointer
	*(&val) = 100
	fmt.Println("local val = ", val)
}

func printGlobal() {
	fmt.Println("Global val=", val)
}

func updateGlobal() {
	val = "updated"
}

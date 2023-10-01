package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {

	// last in first out [stack]
	defer cleanup("first")
	defer cleanup("second")

	fmt.Println("App is working...")

	// defer recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic", err)
		}
	}()

	// panic
	panic("System under attack")
}

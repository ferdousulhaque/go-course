package main

import "fmt"

type author struct {
	name string
}

func main() {
	// declare a map of string keys and author values
	var authors map[string]author

	// initialize the map with make
	authors = make(map[string]author, 0)

	// add authors to the map
	authors["md"] = author{name: "Mohammad"}
	authors["sl"] = author{name: "Sallallahu Alaihi Wa Sallam"}

	fmt.Printf("%#v\n", authors)

	// Alternate
	books := map[string]author{
		"first":  {name: "First book"},
		"second": {name: "Second book"},
	}

	fmt.Printf("%#v\n", books)
	fmt.Println(books["first"])

	// detect if the value was present in map
	a, ok := authors["jm"]
	fmt.Printf("a = %v, ok = %v\n", a, ok)

	books["first"] = author{name: "Testing"}
	fmt.Println(books)

	delete(books, "first")
	fmt.Println(books)

}

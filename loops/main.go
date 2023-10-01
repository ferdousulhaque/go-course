package main

import "fmt"

func main() {
	// For loop on string
	s := "Testing"

	// Loop over
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println("--------------------------")

	// For loop on slice of ints
	sl := []int{100, 200, 300}

	for i, n := range sl {
		fmt.Println(i, n)
	}

	fmt.Println("--------------------------")

	// delcare a map of strings
	ms := map[string]int{
		"us": 110,
		"uk": 95,
		"ca": 100,
	}

	for k, v := range ms {
		fmt.Println(k, v)
	}
}

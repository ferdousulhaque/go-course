package main

import (
	"fmt"
	"runtime"
)

func seperateEvenOdds(n int) (even []int, odd []int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			even = append(even, i)
		} else {
			odd = append(odd, i)
		}
	}
	return even, odd
}

func main() {
	num := 100
	even, odd := seperateEvenOdds(num)
	fmt.Println(even)
	fmt.Println(odd)

	switch os := runtime.GOOS; os {
	case "linux", "darwin", "unix":
		fmt.Println("Linux Variant")
	case "windows":
		fmt.Println("windows variant")
	}

	fmt.Println(runtime.GOOS)
}

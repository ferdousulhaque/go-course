package main

import (
	"fmt"
	"os"
)

func main() {
	// Recovery
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// read file
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %s", err))
	}

	data := string(bs)

	frequency := make(map[string]int)

}

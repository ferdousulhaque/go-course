package main

import (
	"fmt"
	"sync"
)

func main() {
	names := []string{"Ferdous", "Tonni", "Tahan", "Farjad"}

	namesMap := make(map[string]int)

	var mu sync.Mutex

	var wg sync.WaitGroup
	wg.Add(len(names))

	for _, name := range names {
		go func(name string) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			namesMap[name] = len(name)
		}(name)
	}

	wg.Wait()

	fmt.Println(namesMap)
}

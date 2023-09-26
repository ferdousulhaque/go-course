package main

import (
	"fmt"

	"github.com/jboursiquot/go-proverbs"
)

func main() {
	fmt.Println(getProverb())
}

func getProverb() string {
	return proverbs.Random().Saying
}

package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type author struct {
	name string
}

type book struct {
	title string
	author
}

type library map[string][]book

func (l library) addBook(bk book) {
	l[bk.author.name] = append(l[bk.author.name], bk)
}

func (l library) lookupByAuthorName(name string) []book {
	return l[name]
}

func main() {
	lib := library{}

	ub := author{
		name: "Uncle Bob",
	}

	// Add Book
	lib.addBook(book{
		title:  "Domain Driven Design",
		author: ub,
	})

	// Add Book
	lib.addBook(book{
		title:  "Software Architecture",
		author: ub,
	})

	// with spew
	spew.Dump(lib)

	// Search Book by Author name
	found := library.lookupByAuthorName(lib, ub.name)

	fmt.Println(found)
}

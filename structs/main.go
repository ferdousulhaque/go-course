package main

import "fmt"

// struct
type person struct {
	first, last string
}

// attaching function with struct
func (a person) fullName() string {
	return a.first + " " + a.last
}

// attaching a function to change using pointer
func (a *person) changeName(first, last string) {
	a.first = first
	a.last = last
}

type author struct {
	person
	penName string
}

func (a author) fullName() string {
	return a.first + " " + a.last + " [" + a.penName + "]"
}

func main() {
	// struct
	a := person{
		first: "Ferdous",
		last:  "Haque",
	}

	fmt.Printf("%#v\n", a)
	fmt.Println("Fullname " + a.fullName())

	a.changeName("Ferdousul", "Haque")
	fmt.Println("Fullname " + a.fullName())

	b := author{
		person: person{
			first: "A S Md",
			last:  "Ferdousul Haque",
		},
		penName: "Ferdous",
	}

	fmt.Println(b.fullName())
}

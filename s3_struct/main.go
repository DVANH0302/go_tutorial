package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   string
}

func main() {
	var alex person
	alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "abc",
			zip:   "1000",
		},
	}
	alex.print()
	alex.updateName("jimmy")
	alex.print()
}

func (pointerP *person) updateName(newFirstName string) {
	(*pointerP).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v]\n", p)
}

package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (pointer *person) updateNamePointer(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func main() {
	p := person{
		firstName: "Barack",
		lastName:  "Obama",
		contactInfo: contactInfo{
			email:   "obama@usapresident.com",
			zipCode: 9870000,
		},
	}

	fmt.Println(p)
	// p.updateName("ELON MUSK")

	// ppointer := &p
	p.updateNamePointer("ELON")
	p.ijprint2()
}

func (p person) ijprint2() {
	fmt.Printf("%+v", p)
}

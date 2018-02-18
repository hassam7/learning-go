package main

import "fmt"

type contactInfo struct {
	email string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	p := person{
		firstName: "test",
		lastName:  "test two",
		contact: contactInfo{
			email: "test@test.test",
		},
	}
	p.updateName("Jimmy")
	p.showPerson()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
func (p person) showPerson() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	contact ContactInfo
}

type ContactInfo struct {
	email string
	zip int
}

func main(){
	rohan := Person{
		firstName: "Rohan",
		lastName: "Gupta",
		contact: ContactInfo {
		email: "rphan@gmail.com",
		zip: 94000,
		},
	}

	rohan.updateName("jimmy")
	rohan.print()

}


func (p Person) print() {
	fmt.Println("%v",p)
}


func (p *Person)  updateName(newFirstName string){
	(*p).firstName = newFirstName
}

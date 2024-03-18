package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	// contact contactInfo
	contactInfo
}

func main(){
	// alex := person{firstName: "Alex", lastName: "Ardenson"}
	// fmt.Println(alex)

	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// oscar := person{
	// 	firstName: "Oscar",
	// 	lastName: "Oceguera",
	// 	contact: contactInfo{
	// 		email: "oscar@mail.com",
	// 		zipCode: 80140,
	// 	},
	// }
	oscar := person{
		firstName: "Oscar",
		lastName: "Oceguera",
		contactInfo: contactInfo{
			email: "oscar@mail.com",
			zipCode: 80140,
		},
	}
	// fmt.Println(oscar)
	// fmt.Printf("%+v",oscar)
	// oscar.print()

	oscar.updateName("Eduardo")
	oscar.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v",p)

}
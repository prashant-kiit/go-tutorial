package main

import (
	"fmt"
)

type Person struct {
	name    string
	age     uint32
	address Address
}

type Address struct {
	city    string
	country string
}

func (person Person) getFullAddress() string {
	return person.address.city + ", " + person.address.country
}

func main() {
	// default initialization
	var person1 Person
	fmt.Println(person1)

	// explicit initialization
	var person2 Person = Person{
		"Prashant",
		30,
		Address{
			"Bengaluru",
			"India",
		},
	}
	fmt.Println(person2)

	// need to initialize all the members of struct
	// var person3 Person = Person{
	// 	"Prashant",
	// 	30,
	// 	Address{
	// 		"Bengaluru",
	// 	},
	// } // too few values in struct literal of type AddresscompilerInvalidStructLit

	person2.name = "Chinku"
	person2.address.country = "Hindu Rastra"
	fmt.Println(person2)

	// anonymous struct
	person3 := struct {
		name    string
		age     uint32
		address struct {
			city    string
			country string
		}
	}{
		"Mannu",
		20,
		Address{
			"Noida",
			"India",
		},
	}
	fmt.Println(person3)

	// adding method to struct
	fmt.Println(person2.getFullAddress())
}

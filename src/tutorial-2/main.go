package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var age int16 = 30
	var points uint16 = 1000
	// age = age + 400000000; // run time overflow error
	points = points - 4000 // no runtime overflow error

	var cost1 float32 = 12.89788888888 // less accurate
	var cost2 float64 = 12.89788888888 // more accurate
	// var cost2 double = 12.89788888888; // double do not exists

	fmt.Println("Hello Prashant is ", age, points)
	fmt.Println("Costs are", cost1, cost2)

	// fmt.Println("Sum of costs", cost1 + cost2); // type mismatch prevents mathematical operation
	fmt.Println("Sum of costs", float64(cost1)+cost2) // up type casting
	fmt.Println("Sum of costs", float32(cost2)+cost1) // down type casting

	var name string = "Chinku"
	var address string = `Hariom 
	       Megatownship`
	fmt.Println(name + " is from " + address)

	fmt.Println("Length of string in no of bytes", len(name))
	fmt.Println("Runes' representation of a string", []rune(name))
	fmt.Println("Length of Runes' representation of a string", len([]rune(name)))
	fmt.Println("Size of Runes' representation of a string", unsafe.Sizeof([]rune(name)))

	var myRune rune = 'A'
	var myByte byte = 'A'
	fmt.Println("A in rune", myRune)
	fmt.Println("Size of A in rune", unsafe.Sizeof(myRune))
	fmt.Println("A in byte", myByte)
	fmt.Println("Size of A in rune", unsafe.Sizeof(myByte))

	var myBoolean bool = true
	fmt.Println(myBoolean)
	// fmt.Println(int(myBoolean)); // this type casting is not allowed
	// fmt.Println(rune(myBoolean)); // this type casting is not allowed
	// fmt.Println(byte(myBoolean)); // this type casting is not allowed

	var something string // not null
	fmt.Println("Auto initalized string variable", something);
}

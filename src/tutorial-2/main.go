package main

import "fmt"

func main() {
	var age int16 = 30;
	var points uint16 = 1000;
	age = age + 400000000; // run time overflow error
	points = points - 4000; // no runtime overflow error 
	fmt.Println("Hello Prashant is ", age, points);
}
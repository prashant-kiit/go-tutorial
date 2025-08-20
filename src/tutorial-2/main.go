package main

import "fmt"

func main() {
	var age int16 = 30;
	var points uint16 = 1000;
	// age = age + 400000000; // run time overflow error
	points = points - 4000; // no runtime overflow error 

	var cost1 float32 = 12.89788888888; // less accurate
	var cost2 float64 = 12.89788888888; // more accurate
	// var cost2 double = 12.89788888888; // double do not exists

	fmt.Println("Hello Prashant is ", age, points);
	fmt.Println(cost1, cost2);
}
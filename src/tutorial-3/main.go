package main

import "fmt"

func main() {
	var result float64 = divide(11, 2)
	fmt.Println(result)
}

func divide(numerator float64, denominator float64) float64 {
	var result float64 = numerator / denominator;
	return result
}

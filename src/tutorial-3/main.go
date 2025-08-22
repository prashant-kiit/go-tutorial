package main

import "fmt"

func main() {
	var result float64 = divide(11, 2)
	var quotient, rem int = divideByInt(11, 2)
	fmt.Println(result)
	fmt.Println(quotient, rem)
}

func divide(numerator float64, denominator float64) float64 {
	var result float64 = numerator / denominator
	return result
}

func divideByInt(numerator int, denominator int) (int, int) {
	var quotient int = numerator / denominator
	var rem int = numerator % denominator
	return quotient, rem
}

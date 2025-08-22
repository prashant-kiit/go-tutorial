package main

import (
	"errors"
	"fmt"
)

func main() {
	var result, err1 = divide(11, 0)
	var quotient, rem, err2 = divideByInt(11, 2)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	fmt.Println(result)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	fmt.Println(quotient, rem)
}

func divide(numerator float64, denominator float64) (float64, error) {
	var err error
	if denominator == 0 {
		err = errors.New("the denominator is zero")
		return 0, err
	}
	var result float64 = numerator / denominator
	return result, err
}

func divideByInt(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("the denominator is zero")
		return 0, 0, err
	}
	var quotient int = numerator / denominator
	var rem int = numerator % denominator
	return quotient, rem, err
}

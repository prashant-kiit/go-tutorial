package main

import "fmt"

// Generic struct with type parameter T
type Box[T any] struct {
	Value T
}

// Generic method
func (b Box[T]) GetValue() T {
	return b.Value
}

func main() {
	intBox := Box[int]{Value: 42}
	strBox := Box[string]{Value: "Hello"}

	fmt.Println("Int Box:", intBox.GetValue())
	fmt.Println("String Box:", strBox.GetValue())
}

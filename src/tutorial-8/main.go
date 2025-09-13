package main

import "fmt"

// function to show pointer update
func update(val *int) {
	*val = *val + 100
}

func main() {
	// 1. Basic Pointer
	var a int = 10
	var p *int = &a
	fmt.Println("=== Basic Pointer ===")
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p:", p)
	fmt.Println("Value at *p:", *p)

	// 2. Changing Value Through Pointer
	fmt.Println("\n=== Modify via Pointer ===")
	b := 20
	q := &b
	fmt.Println("Before:", b)
	*q = 50
	fmt.Println("After:", b)

	// 3. Pointer Passed to Function
	fmt.Println("\n=== Pointer to Function ===")
	c := 5
	fmt.Println("Before update:", c)
	update(&c)
	fmt.Println("After update:", c)

	// 4. new() and Pointer
	fmt.Println("\n=== new() and Pointer ===")
	r := new(int) // allocates memory, initializes to 0
	fmt.Println("Default:", *r)
	*r = 42
	fmt.Println("Updated:", *r)
}

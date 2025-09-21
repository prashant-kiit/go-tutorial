package main


import "fmt"

// Generic function to print a slice of any type
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func main() {
	ints := []int{1, 2, 3}
	strs := []string{"Go", "Generics", "Example"}

	fmt.Println("Int slice:")
	PrintSlice(ints)

	fmt.Println("\nString slice:")
	PrintSlice(strs)
}

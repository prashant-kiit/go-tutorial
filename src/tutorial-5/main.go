package main

import "fmt"

func main() {
	var name string = "Präshänt"
	for i, v := range name {
		fmt.Println(i, v)
	}

	var name1 []byte = []byte("Präshänt")
	for i, v := range name1 {
		fmt.Println(i, v)
	}

	var name2 []rune = []rune("Präshänt")
	for i, v := range name2 {
		fmt.Println(i, v)
	}
}

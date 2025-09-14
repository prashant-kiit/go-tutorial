package main

import (
	"fmt"
	"github.com/prashant-kiit/go-tutorial/src/tutorial-9/myPkg"
)

// packages/modules
// import & export
// accessiblity scope (public & private)

func main() {
	fmt.Println("Hello, World!")
	// Create Person (can only set exported fields)
    p := myPkg.Person{Name: "Prashant"} // ✅ OK
	fmt.Println("Name:", p.Name)

	// p.age = 30 // ❌ not allowed (unexported)

	// Call exported function
    myPkg.SayHello(p)

	// mypkg.secretMessage() // ❌ not allowed (private)
}
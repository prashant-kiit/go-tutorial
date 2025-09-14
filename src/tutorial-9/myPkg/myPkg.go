package myPkg

import "fmt"

// Exported (public) struct
type Person struct {
    Name string // exported
    age  int    // private
}

// Exported (public) function
func SayHello(p Person) {
    fmt.Println("Hello,", p.Name)
	p.secretMessage() // OK here (same package)
}

// Unexported (private) function
func (p Person) secretMessage() {
    fmt.Println("This is private to mypkg")
	fmt.Println("Age:", p.age)
}


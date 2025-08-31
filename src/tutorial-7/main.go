package main

import "fmt"

type ModernEngine struct {
	mileage uint32
}

type ClassicEngine struct {
	mileage uint32
}

type Engine interface {
	getMileage() uint32
}

func (e ModernEngine) getMileage() uint32 {
	return e.mileage
}

func (e ClassicEngine) getMileage() uint32 {
	return e.mileage/2
}

func main() {
	var engine1 Engine = ModernEngine{100}
	var engine2 Engine = ClassicEngine{50}
	fmt.Println(engine1.getMileage())
	fmt.Println(engine2.getMileage())
}

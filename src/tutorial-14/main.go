package main

import (
	"fmt"
	"time"
)

func foo1(ch chan int, length int) {
	for i := range length + 10{ // no blocking
		ch <- i
		time.Sleep(300 * time.Millisecond)
	}
}

func foo2(ch chan int, length int) {
	for i := range length + 10{ // no blocking
		ch <- i
		time.Sleep(30000 * time.Millisecond)
	}
}

func example1() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	ch <- 3

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 4
	fmt.Println(<-ch)
}

func example2() {
	length := 10
	ch := make(chan int, length)
	for i := range length {
		ch <- i
	}
	for range length {
		fmt.Println(<-ch)
	}
}

func example3() {
	length := 10
	ch := make(chan int, length)

	go foo1(ch, length)

	for range length {
		fmt.Println(<-ch)
	}
}

func exmaple4() {
	length := 10
	ch := make(chan int, length)

	go foo1(ch, length)
	go foo1(ch, length)

	for range length {
		fmt.Println(<-ch)
	}
}

func exmaple5() {
	length := 10
	ch1 := make(chan int, length)
	ch2 := make(chan int, length)

	go foo1(ch1, length)
	go foo2(ch2, length)

	for range length {
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
	}
}

func exmaple6() {
	length := 10
	ch1 := make(chan int, length)
	ch2 := make(chan int, length)

	go foo1(ch1, length)
	go foo2(ch2, length)

	for range length {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	// example1()
	// example2()
	// example3()
	// exmaple4()
	// exmaple5()
	exmaple6()
}

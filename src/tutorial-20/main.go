package main

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() { ch1 <- "from ch1" }()
    go func() { ch2 <- "from ch2" }()

    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
// Go code must be inside main()
// select randomly picks one ready channel
// Use loop/buffer/done/context to avoid leaks
// Any one of the channels will be in blocking state
// At the end both channels should be closed to avoid leaks
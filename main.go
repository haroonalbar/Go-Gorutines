package main

import (
	"fmt"
	"time"
)

func say(s string, i int) {
	fmt.Printf("%v: %v\n", s, i)
}

func main() {
	for i := 0; i < 5; i++ {
		go say("Hello", i)
		go say("Boi", i)
	}
	time.Sleep(1 * time.Millisecond)
}

// here im seeing a pattern where all the go routines are not executed on 30 ms
// execution pause there both gorutines are run but the result im getting is only upto i=2
// and if I give more execution pause more of the function is executed.
// i'll look further into it later

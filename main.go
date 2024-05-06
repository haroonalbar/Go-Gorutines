package main

import (
	"fmt"
	"time"
)

func say (s string) {
  for i:=0 ; i < 5 ; i++{
    time.Sleep(10*time.Millisecond)
    fmt.Printf("%v: %v\n",s,i)
  }
}

func main() {
  go say("Hello")
  go say("Boi")
  time.Sleep(30*time.Millisecond)
}
// here im seeing a pattern where all the go routines are not executed on 30 ms 
// execution pause there both gorutines are run but the result im getting is only upto i=2 
// and if I give more execution pause more of the function is executed.
// i'll look further into it later

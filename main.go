package main

import "fmt"

func say (s string) {
  for i:=0 ; i < 5 ; i++{
    fmt.Println(s)
  }
}

func main() {
  say("Hello")
  say("Boi")
}

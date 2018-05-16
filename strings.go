package main

import (
  "fmt"
  "strings"
)

func main() {
  greetings := [] string{"Hello", "world!"}
  newString := strings.Join(greetings, " ")
  fmt.Println(newString)
  fmt.Printf("Length of the new string: %d\n", len(newString))

}

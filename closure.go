package main

import (
  "fmt"
)

func getSequence() func() int {
  i:=0
  return func() int {
    i+=1
    return i
  }
}

func main() {
  /* nextNumber is now a function with i as 0 */
  nextNumber := getSequence()

  fmt.Println(nextNumber())
  fmt.Println(nextNumber())
}

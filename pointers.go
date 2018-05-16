package main

import "fmt"

func main() {
  var a int = 10
  fmt.Printf("Address of a variable %x\n", &a)

  var ip *int /* pointer to the variable */

  ip = &a

  fmt.Printf("Address of a variable: %x\n", ip)

  /* access the value stored in the pointer */
  fmt.Printf("Value of the pointer: %d\n", *ip)
}

package main

import "fmt"

func main() {
  numbers := []int{0,1,2,3,4,5,6,7,8}
  printSlice(numbers)

  /* print the original numbers */
  fmt.Println("number == ", numbers)

  /* print the subslice */
  fmt.Println("numbers[1:4] == ", numbers[1:4])

  /* print the lower bound of the slice */
  fmt.Println("numbers[:3] == ", numbers[:3])

  /* append new number */
  numbers = append(numbers, 999)

  printSlice(numbers)

  fmt.Println("New slice is: ", numbers)
}

func printSlice(x [] int) {
  fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

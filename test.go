package main

import "fmt"

func main() {
  var j = 3
  var x float64
  x = 20.0
  j = j + 1
  y := 54
  fmt.Printf("y is of type %T\n", y)
  fmt.Printf("x is of type %T\n", x)

  if (j < 20) {
    /* j is less than 20 */
    fmt.Printf("j is less than 20, and its value is %d\n", j)
  } else {
    fmt.Printf("j is more than 20")
  }

  /* switch */
  var marks int = 90
  var grade string = "B"
  switch marks {
    case 90: grade = "A"
    case 80: grade = "B"
    case 50,60,70: grade = "C"
    default: grade = "D"
  }
  fmt.Printf("your grade is %s\n", grade)

  switch {
  case grade == "A":
    fmt.Printf("Excellent\n")
  case grade == "B":
      fmt.Printf("Well done\n")
  }

  /* loops */
  var b int = 15
  var a int
  numbers := [6]int{1,2,3,4,5}

  for a=0; a <10; a++ {
    fmt.Printf("Value of a is %d\n", a)
  }

  for a < b {
    a++
    fmt.Printf("value of a: %d\n", a)
  }

  for i, x:=range numbers {
    fmt.Printf("value of x = %d at %d\n", x, i)
  }

  fmt.Printf("division of 4/3=%d\n",4/3)

  printPrimeNumbers()

  var firstNum, secondNum = 3,5
  fmt.Printf("Biggest number is: %d\n", max(firstNum, secondNum))

  /* print my name in reverse */
  var firstName, secondName = "Wojtek", "Hejna"
  firstName, secondName = swap(firstName, secondName)
  fmt.Printf("My name in reverse: %s %s\n", firstName, secondName)
}

func printPrimeNumbers() {
  var i, j int

  for i = 2; i < 100; i++ {
    for j = 2; j < (i/j); j++ {
      if (i%j == 0) {
        break; // if factor found, not prime
      }
    }
    if (j > (i/j)) {
      fmt.Printf("%d is prime\n", i)
    }
  }
}

func max(num1, num2 int) int {
  var result int

  if (num1 > num2) {
    result = num1
  } else {
    result = num2
  }

  return result
}

func swap(x, y string) (string, string) {
  return y, x
}

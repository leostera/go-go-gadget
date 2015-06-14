package main

import(
  "fmt"
)

func add (x, y int) (result int) {
  result = x + y
  return
}

func main() {
  fmt.Print(add(5,9))
}


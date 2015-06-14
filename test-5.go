package main

import(
  "fmt"
)

func add (x, y int) (int, int, int) {
  return x, y, x + y
}

func main() {
  x, y, result := add(3, 4)
  fmt.Print(x, y, result)
}


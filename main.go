package main

import (
  "fmt"
)

func main() {
  arg := 1000
  fmt.Println(some(arg))
}

func some[T any](arg T) T {
  return arg
}

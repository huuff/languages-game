package main

import (
  "fmt"
  "os"
)

func main() {
  name := os.Args[1]
  result := fmt.Sprintf("Hello %s!", name)
  fmt.Println(result)
}

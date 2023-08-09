package main

import (
  "fmt"
  "os"
)

func main() {
  if len(os.Args) > 1 {
    name := os.Args[1]
    result := fmt.Sprintf("Hello %s!", name)
    fmt.Println(result)
  } else {
    fmt.Println("Hello World!")
  }
}

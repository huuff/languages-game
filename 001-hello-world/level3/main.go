package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  if len(os.Args) > 1 {
    name := os.Args[1]

    currentTime := time.Now()
    // TODO: A function to create todays date with different times
    lastMorningTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 12, 0, 0, 0, currentTime.Location())
    lastAfternoonTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 18, 0, 0, 0, currentTime.Location())

    // TODO: Use a switch
    // TODO: Maybe use a Greeter interface for this?
    if currentTime.Before(lastMorningTime) {
      fmt.Println(fmt.Sprintf("Good morning, %s!", name))
    } else if currentTime.Before(lastAfternoonTime) {
      fmt.Println(fmt.Sprintf("Good afternoon, %s!", name))
    } else {
      fmt.Println(fmt.Sprintf("Good evening, %s!", name))
    }

  } else {
    fmt.Println("Hello World!")
  }
}

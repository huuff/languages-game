package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	switch {
	case input%15 == 0:
		fmt.Println("FizzBuzz")
	case input%3 == 0:
		fmt.Println("Fizz")
	case input%5 == 0:
		fmt.Println("Buzz")
	default:
		fmt.Println(input)
	}

}

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var delay int
	var err error
	if len(os.Args) > 1 {
		delay, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Wrong rate: %s", os.Args[1])
		}
	} else {
		delay = 1000
	}

	ticker := time.NewTicker(time.Duration(delay) * time.Millisecond)

	tickerChannel := ticker.C
	numberChannel := make(chan int)

	go func() {
		for num := 0; true; num++ {
			<-tickerChannel
			numberChannel <- num
		}
	}()

	go func() {
		for {
			num := <-numberChannel
			switch {
			case num%15 == 0:
				fmt.Println("FizzBuzz")
			case num%3 == 0:
				fmt.Println("Fizz")
			case num%5 == 0:
				fmt.Println("Buzz")
			default:
				fmt.Println(num)
			}
		}
	}()

	for {
		time.Sleep(time.Second)
	}
}

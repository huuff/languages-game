package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

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

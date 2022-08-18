package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)

	go func() {
		defer fmt.Println("go channel end")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("len(c)", len(c), "cap(c)", cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-c

		fmt.Println("num = ", num)
	}

	fmt.Println("main end")
}

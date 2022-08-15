package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("defer end")
		fmt.Println("gorutine is run")
		c <- 888
	}()

	num := <-c
	fmt.Println("main run", num)
}

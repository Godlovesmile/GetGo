package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("test")

	go func(a int, b int) bool {
		fmt.Println("a, b", a, b)
		return true
	}(10, 10)

	go func() {
		defer fmt.Println("A defer")

		func() {
			defer fmt.Println("B defer")
			runtime.Goexit()
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	// loop time
	for {
		time.Sleep(1 * time.Second)
	}
}

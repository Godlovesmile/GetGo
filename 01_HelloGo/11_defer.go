package main

import "fmt"

// defer 类似压栈机制
// main end; defer run2; defer run1
func main() {
	defer fmt.Println("defer run1")
	defer fmt.Println("defer run2")

	fmt.Println("main end")
}

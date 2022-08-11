package main

import "fmt"

func main() {
	fmt.Printf("Hello, world or 你好世界\n")

	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
	fmt.Printf("type of s = %T", s)
}

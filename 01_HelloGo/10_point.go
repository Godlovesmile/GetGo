package main

import "fmt"

func swap(pa *int, pb *int) {
	temp := *pa
	*pa = *pb
	*pb = temp
}

func main() {
	a, b := 10, 20
	swap(&a, &b)

	fmt.Println(a, b)
}

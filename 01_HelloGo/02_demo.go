package main

// import "fmt"

// type testInt func(int) bool // 声明一个函数类型

// func isOdd(i int) bool {
// 	if i%2 == 0 {
// 		return false
// 	}
// 	return true
// }

// // 声明的函数类型当做一个参数
// func filter(temparr []int, f testInt) []int {
// 	var res []int
// 	for _, val := range temparr {
// 		if f(val) {
// 			res = append(res, val)
// 		}
// 	}
// 	return res
// }

// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	odd := filter(slice, isOdd)
// 	fmt.Println("Odd elements are: ", odd)
// }

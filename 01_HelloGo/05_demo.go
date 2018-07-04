package main

// import (
// 	"fmt"
// )

// // Human is
// type Human struct {
// 	name  string
// 	age   int
// 	phone string
// }

// type Student struct {
// 	Human  // 匿名字段
// 	school string
// 	loan   float32
// }

// type Employee struct {
// 	Human   // 匿名字段
// 	company string
// 	money   float32
// }

// // Huamn 实现 SayHi 方法
// func (h Human) SayHi() {
// 	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
// }

// // Human 实现 Sing 方法
// func (h Human) Sing(lyrics string) {
// 	fmt.Printf("la la la...%s\n", lyrics)
// }

// // Employee 重载 Human 的 SayHi 方法
// func (e Employee) SayHi() {
// 	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
// }

// // Interface Men 被 Human, Student, Employee 实现
// // 因为这三类都实现了这两个方法
// type Men interface {
// 	SayHi()
// 	Sing(lyrics string)
// }

// func main() {
// 	mike := Student{Human{"Mike", 25, "15501251520"}, "MIT", 0.00}

// 	// 定义 Men 类型的变量 i
// 	var i Men
// 	// i 能存储 Student
// 	i = mike
// 	fmt.Printf("this is a student, is %s\n", mike.name)
// 	i.SayHi()
// 	i.Sing("er quan yin yue")
// }

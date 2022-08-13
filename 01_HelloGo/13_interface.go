package main

import "fmt"

// ==== 1. interface 多态 ====
type Animal interface {
	GetName() string
	Sleep()
}

type Cat struct {
	name string
}

func (self *Cat) GetName() string {
	return self.name
}

func (self *Cat) Sleep() {
	fmt.Println("cat sleep")
}

type Dog struct {
	name string
}

func (self *Dog) GetName() string {
	return self.name
}

func (self *Dog) Sleep() {
	fmt.Println("dog sleep")
}

func showAnimal(animal Animal) {
	animal.Sleep()
}

// ===== 2. interface ====实现万能数据类型
func Test(arg interface{}) {
	// 通过 类型断言机制 判断 arg 类型
	_, ok := arg.(string)

	if ok {
		fmt.Println("arg is string")
	}
}

func main() {
	c := Cat{"cat type"}
	d := Dog{"dog type"}

	showAnimal(&c)
	showAnimal(&d)

	Test("hahah")
	Test(888)
}

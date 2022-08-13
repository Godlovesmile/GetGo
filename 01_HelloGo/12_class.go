package main

import "fmt"

type Man struct {
	name string
	age  int
}

type SuperMan struct {
	Man // SuperMan 继承 Man

	height float32
}

func (self *Man) GetManName() string {
	return self.name
}

func (self *Man) SetManAge(age int) {
	self.age = age
}

func main() {
	m := Man{"ylz", 20}
	fmt.Println("name = ", m.GetManName())
	m.SetManAge(18)
	fmt.Println("age = ", m.age)

	s := SuperMan{Man{"xiao", 33}, 188}
	fmt.Println("xiao name = ", s.name)
	fmt.Println("xiao height = ", s.height)
}

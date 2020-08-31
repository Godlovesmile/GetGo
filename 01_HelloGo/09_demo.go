package main

import "fmt"

// Human is a representation of a person
type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human
	scholl string
}

type Elpee struct {
	Human
	company string
}

func (h *Human) sayHi() {
	fmt.Printf("say hi\n")
}

func (s *Student) sayHi() {
	fmt.Printf("student say hi\n")
}

func main() {
	s1 := Student{Human{"ylz", "155000000000", 18}, "jializudn"}
	fmt.Printf("name--->>>%s\n", s1.name)
	s1.sayHi()
}

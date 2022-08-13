package main

import "fmt"

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

func main() {
	c := Cat{"cat type"}
	d := Dog{"dog type"}

	showAnimal(&c)
	showAnimal(&d)
}

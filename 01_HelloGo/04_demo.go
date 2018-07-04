package main

// import (
// 	"fmt"
// )

// const (
// 	WHITE = iota
// 	BLACK
// 	BLUE
// 	RED
// 	YELLOW
// )

// type Color byte

// // Box is a representation of a Box
// type Box struct {
// 	width, height, depth float64
// 	color                Color
// }

// // BoxList is box
// type BoxList []Box

// func (b Box) Volume() float64 {
// 	return b.width * b.height * b.depth
// }

// func (b *Box) SetColor(c Color) {
// 	b.color = c
// }

// func (b1 BoxList) BiggestsColor() Color {
// 	v := 0.00
// 	k := Color(WHITE)
// 	for _, b := range b1 {
// 		if b.Volume() > v {
// 			v = b.Volume()
// 			k = b.color
// 		}
// 	}
// 	return k
// }

// func (bl BoxList) PaintItBlock() {
// 	for i, _ := range bl {
// 		bl[i].SetColor(BLACK)
// 	}
// }

// func (c Color) String() string {
// 	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
// 	return strings[c]
// }

// func main() {
// 	boxes := BoxList{
// 		Box{4, 4, 4, RED},
// 		Box{10, 10, 1, YELLOW},
// 		Box{1, 1, 20, BLACK},
// 		Box{10, 20, 20, BLUE},
// 	}
// 	fmt.Printf("We have %d boxes in our set\n", len(boxes))
// 	fmt.Println("I want to change bigger box color")
// 	boxes[3].SetColor(BLACK)
// 	fmt.Printf("The biggest one is %s", boxes.BiggestsColor().String())
// }

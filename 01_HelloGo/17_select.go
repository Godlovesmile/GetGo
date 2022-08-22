package main

import (
	"fmt"
	"time"
)

/*
select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。

select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
*/
func Chann(c chan int, stopC chan bool) {
	for i := 0; i < 10; i++ {
		c <- i

		time.Sleep(time.Second)
	}

	stopC <- true
}

func main() {
	c := make(chan int)
	stopChann := make(chan bool)

	go Chann(c, stopChann)

	for {
		select {
		case g := <-c:
			fmt.Println("receive g", g)
		case s := <-c:
			fmt.Println("receive s", s)
		case _ = <-stopChann:
			goto end
		}
	}
end:
}

package main

import (
	"fmt"
)

func count_chan(ch chan int) {
	fmt.Println("Counting")
	ch <- 1
}

func main() {
	// chs := make([]chan int, 10)

	// for i := 0; i < 10; i++ {
	// 	chs[i] = make(chan int)
	// 	go count_chan(chs[i])
	// }
	// for _, v := range chs {
	// 	<-v
	// }
	// 	ch := make(chan int, 1)
	// 	for {
	// 		select {
	// 		case ch <- 0:
	// 		case ch <- 1:
	// 		}
	// 		i := <-ch
	// 		fmt.Printf("i: %v\n", i)
	// 	}

	// ch := make(chan int, 1)

	// timeout := make(chan bool, 1)
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	timeout <- true
	// }()

	// select {
	// case <-ch:
	// 	i := <-ch
	// 	fmt.Printf("i: %v\n", i)
	// case <-timeout:
	// 	fmt.Println("timeout")
	// }

}

// func parse(ch <-chan int) {
// 	for v := range ch {
// 		fmt.Printf("v: %v\n", v)
// 	}
// 	ch2 := chan int(ch)
// 	ch2 <- 2
// 	for v := range ch2 {
// 		fmt.Printf("v: %v\n", v)
// 	}
// }

type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

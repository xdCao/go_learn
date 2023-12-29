package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go Add(i, i)
	}
	time.Sleep(2 * time.Second)
}

func Add(x, y int) {
	z := x + y
	fmt.Printf("z: %v\n", z)
}

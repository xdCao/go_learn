package main

import (
	"fmt"
	"sync"
	"time"
)

var a string
var once sync.Once

func setup() {
	a = "hello, world"
	fmt.Println("setup")
}
func doprint() {
	once.Do(setup)
	print(a)
}
func twoprint() {
	go doprint()
	go doprint()
}

func main() {
	twoprint()
	time.Sleep(1 * time.Second)
}

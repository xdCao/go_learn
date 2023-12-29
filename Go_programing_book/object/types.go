package main

import "fmt"

func main() {
	// 	var a Integer = 1
	// 	a.Add(2)
	// 	fmt.Println(a)

	a := [3]int{1, 2, 3}
	b := a
	b[0]++
	fmt.Println(a, b) // [1 2 3] [2 2 3]
	b2 := &a
	b2[0]++
	fmt.Println(a, *b2) // [2 2 3] [2 2 3]

}

type Integer int

func (this *Integer) Add(input Integer) {
	*this += input
}

// func (this Integer) Add(input Integer) {
// 	this += input
// }

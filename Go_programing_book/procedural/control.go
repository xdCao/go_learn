package main

import "fmt"

func main() {
	// fmt.Println(if_func(0))
	// switch_func(2)
	// switch_func(4)
	// switch_func(7)
	for_func()
}

func if_func(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func switch_func(i int) {
	switch i {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough
	case 3:
		fmt.Println("3")
	case 4, 5, 6:
		fmt.Println("456")
	default:
		fmt.Println("default")
	}

	switch {
	case 0 <= i && i <= 3:
		fmt.Println("0-3")
	case 4 <= i && i <= 6:
		fmt.Println("4-6")
	}
}

func for_func() {
	a := []int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
}

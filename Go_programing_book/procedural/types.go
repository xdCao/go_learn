package main

import "fmt"

func main() {
	// array()
	// slice()
	map_func()
}

// 数组例程
func array() {
	arr := [5]int{1, 2, 3, 4, 5}
	modify(arr)
	fmt.Println("main arr = ", arr)

	for i := 0; i < len(arr); i++ {

	}

	for idx, v := range arr {
		fmt.Println(idx, v)
	}
}

func modify(arr [5]int) {
	arr[0] = 10
	fmt.Println("modify arr = ", arr)
}

func slice() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[:5]
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Println(len(slice))
	slice = append(slice, 6)
	fmt.Println(cap(slice))

	slice1 := make([]int, 5)
	slice2 := make([]int, 5, 10)
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice1 = append(slice1, 2, 3, 4)
	slice1 = append(slice1, slice1...)
	slice1 = append(slice1, slice2...)

	fmt.Println(slice1)

	slice1 = []int{1, 2, 3, 4, 5}
	slice2 = []int{6, 7, 8}
	copy(slice2, slice1)
	fmt.Println(slice2)
	// [1 2 3]
	slice2 = []int{6, 7, 8}
	copy(slice1, slice2)
	fmt.Println(slice1)
	// [6 7 8 4 5]
}

func map_func() {
	personDB := make(map[string]Person)
	personDB["1"] = Person{ID: "1", Name: "Tom", Address: "Beijing"}
	personDB["2"] = Person{ID: "2", Name: "Jack", Address: "Shanghai"}

	target, ok := personDB["3"]
	if ok {
		fmt.Println(target)
	} else {
		fmt.Println("person not found")
	}

	target, ok = personDB["2"]
	if ok {
		fmt.Println(target)
	} else {
		fmt.Println("person not found")
	}
}

type Person struct {
	ID      string
	Name    string
	Address string
}

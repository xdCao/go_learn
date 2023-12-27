package main

import (
	"errors"
	"fmt"
)

func main() {
	// c, ok := Add(-1, -2)
	// if ok == nil {
	// 	fmt.Print(c)
	// } else {
	// 	fmt.Print(ok)
	// }
	// args_func(1, 2, 3)
	// type_print(1, 2, "s", 1.234)

	// anoy_func()

	closure_func()

}

func Add(a, b int) (ret int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("只支持正数加法")
		return
	}
	return a + b, nil
}

func args_func(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
	args_func(args[:1]...)
}

func type_print(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func anoy_func() {
	f := func(x, y int) int {
		return x + y
	}
	res := f(1, 2)
	fmt.Println(res)
}

func closure_func() {

	j := 5
	a := func() func() {
		i := 10
		return func() {
			fmt.Println("i = ", i, ", j = ", j)
		}
	}()

	a()
	j *= 2
	a()

}

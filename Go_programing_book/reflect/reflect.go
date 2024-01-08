package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var x float64 = 3.14
	// fmt.Printf("x: %T\n", x)
	// fmt.Println("type of x : ", reflect.TypeOf(x))
	// v := reflect.ValueOf(x)
	// fmt.Printf("v type : %v\n", v.Type())
	// fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	// fmt.Println("value: ", v.Float())

	// var x float64 = 3.4
	// p := reflect.ValueOf(&x) // 注意:得到X的地址
	// fmt.Println("type of p:", p.Type())
	// fmt.Println("settability of p:", p.CanSet())
	// v := p.Elem()
	// fmt.Println("type of v:", v.Type())
	// fmt.Println("settability of v:", v.CanSet())
	// v.SetFloat(7.1)
	// fmt.Println(v.Interface())
	// fmt.Println(x)

	t := T{203, "mh203"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

type T struct {
	A int
	B string
}

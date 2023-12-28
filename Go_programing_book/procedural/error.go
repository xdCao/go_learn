package main

import "fmt"

func main() {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("runtime error caught, %v", r)
		}
	}()

	foo()

}

func foo() {
	panic("oops")
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string {
	return e.Op + " " + e.Path + ": " + e.Err.Error()
}

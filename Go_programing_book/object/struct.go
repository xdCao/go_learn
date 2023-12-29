package main

import "fmt"

func main() {

	// rect1 := new(Rect)
	// rect1 = &Rect{}
	// rect1 = &Rect{0, 0, 100, 200}
	// rect1 = &Rect{width: 100, height: 200}
	// rect1 = NewRect(0, 0, 100, 200)

	// fmt.Println(rect1)

	base := &Base{Name: "xxx"}
	base.Foo() //

	foo := &Foo{Base: base, Alias: "alias"}
	foo.Foo() // xxx \n alias \n

}

type Rect struct {
	x, y          float64
	width, height float64
}

func (this *Rect) Area() float64 {
	return this.width * this.height
}

func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

type Base struct {
	Name string
}

func (this *Base) Foo() {
	fmt.Println(this.Name)
}

type Foo struct {
	*Base
	Alias string
}

func (this *Foo) Foo() {
	this.Base.Foo()
	fmt.Println(this.Alias)
}

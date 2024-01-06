package main

import (
	"encoding/json"
	"fmt"
)

type TestJson struct {
	Title     string
	Authors   []string
	Published bool
	Price     float32
	Sales     int
}

func main() {
	obj := &TestJson{
		Title:     "Go语言编程",
		Authors:   []string{"111", "222"},
		Published: true,
		Price:     9.99,
		Sales:     100000,
	}
	bytes, _ := json.Marshal(obj)
	jsonstr := string(bytes)
	fmt.Printf("jsonstr: %v\n", jsonstr)

	var container interface{}
	json.Unmarshal([]byte(jsonstr), &container)

	mapobj, ok := container.(map[string]interface{})
	if ok {
		for k, v := range mapobj {
			fmt.Printf("k: %s, v : %v, type : %T\n", k, v, v)
		}
	}

}

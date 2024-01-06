package main

import (
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":9999", nil)

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	test_get()
}

func test_get() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}

func test_post() {
	file, _ := os.Open("pic.png")
	reader := bufio.NewReader(file)
	resp, err := http.Post("http://example.com/upload", "image/jpeg", reader)
	if err != nil {
		// 处理错误
		return
	}
	if resp.StatusCode != http.StatusOK {
		// 处理错误
		return
	}
	// ...
}

func test_post_form() {
	resp, err := http.PostForm(
		"http://example.com/posts",
		url.Values{
			"title":   {"article title"},
			"content": {"article body"},
		},
	)
	if err != nil { // 处理错误
		return
	}
	fmt.Printf("resp: %v\n", resp)
	// ...
}

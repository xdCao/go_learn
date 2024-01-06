package main

import (
	"fmt"
	"net/http"
)

type CustomTransport struct {
	Transport http.RoundTripper
}

func (t *CustomTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// do something
	return t.transport().RoundTrip(req)
}

func (t *CustomTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func main() {
	t := &CustomTransport{}
	c := t.Client()
	resp, err := c.Get("http://example.com")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("resp: %v\n", resp)
}

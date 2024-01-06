package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	enc := json.NewEncoder(os.Stdout)

	for {
		var v map[string]interface{}
		err := dec.Decode(&v)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		for k, _ := range v {
			if k != "title" {
				v[k] = nil
			}
		}
		err = enc.Encode(&v)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
}

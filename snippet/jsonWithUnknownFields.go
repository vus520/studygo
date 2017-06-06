package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	A int                    `json:"a"`
	B int                    `json:"b"`
	X map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

func main() {
	s := `{"a":1, "b":2, "x":1, "y":1}`
	f := Foo{}
	if err := json.Unmarshal([]byte(s), &f.X); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n\n", f)

	if n, ok := f.X["a"].(float64); ok {
		f.A = int(n)
	}
	if n, ok := f.X["b"].(float64); ok {
		f.B = int(n)
	}
	delete(f.X, "a")
	delete(f.X, "b")

	fmt.Printf("%+v", f)
}

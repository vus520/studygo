package main

import "fmt"

func main() {

	var a interface{} = "i'm string"

	switch a.(type) {
	case string:
		fmt.Println("interface a's type is string")
	}
}

package main

import "fmt"

func main() {

	a := []int{1, 2, 3}

	fmt.Println(a)

	a = a[:len(a)-1]
	fmt.Println(a)

	a = a[1:]
	fmt.Println(a)
}

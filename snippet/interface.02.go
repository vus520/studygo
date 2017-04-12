package main

import "fmt"

func main() {
	var t interface{}
	t = 1
	t0 := (t).(int)
	fmt.Printf("%v\t%v\n", t0, &t0)

	t = "ok"
	t1 := (t).(string)
	fmt.Printf("%v\t%v\n", t1, &t1)
}
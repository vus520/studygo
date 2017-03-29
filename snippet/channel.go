package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Millisecond * time.Duration(sec))

	fmt.Println(w, "is ready")

	c <- sec
}

func main() {

	c = make(chan int)

	go ready("Tee", 2)
	go ready("Coffee", 1)

	fmt.Println("I'm waiting")

	x, y := <-c, <-c

	fmt.Printf("x=%s, y=%s", x, y)
}

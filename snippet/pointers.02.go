package main

import (
	"fmt"
	"strings"
)

type User struct {
	Id   int
	Name string
}

func test0() {

	up := &User{0, "nobody"}
	up.Id = 1
	up.Name = "Jack"

	fmt.Println("第一次值", up)
	u2 := up
	u2.Name = "Tom"
	fmt.Println("修改后原来的值", up)
	fmt.Println("修改的值", u2)

}


func test1() {

	up := &User{0, "nobody"}
	up.Id = 1
	up.Name = "Jack"

	fmt.Println("第一次值", up)
	u2 := *up
	u2.Name = "Tom"
	fmt.Println("修改后原来的值", up)
	fmt.Println("修改的值", u2)

}

func test2() {

	up := User{0, "nobody"}
	up.Id = 1
	up.Name = "Jack"

	fmt.Println("第一次值", up)
	u2 := &up
	u2.Name = "Tom"
	fmt.Println("修改后原来的值", up)
	fmt.Println("修改的值", u2)

}

func main() {
	fmt.Println(strings.Repeat("-", 50))
	test0()
	fmt.Println(strings.Repeat("-", 50))
	test1()
	fmt.Println(strings.Repeat("-", 50))
	test2()
}
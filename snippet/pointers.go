package main

import "fmt"


func main() {

    s  := "this first str"
    ss := s;

    fmt.Println("ss:" + ss)
    s  = "this secend str"
    fmt.Println("after modify ss:" + ss)

    aa := &s
    s  = "this third str"
    fmt.Println("aa = &s; *ss:" + *aa)


    i := 1
    fmt.Println("initial:", i)
	zeroval(i)
    fmt.Println("zeroval:", i)

    zeroptr(&i)
    fmt.Println("zeroptr:", i)

	zeroval(i)
    fmt.Println("zeroval:", i)

    fmt.Println("pointer:", &i)
}


func zeroval(ival int) {
    ival = 0
}

func zeroptr(iptr *int) {
    *iptr = 0
}
package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "log"
)

type P struct {
    X, Y, Z int
    Name    string
}

type Q struct {
    X, Y *int32
    Name string
}

func main() {
    var network bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&network) // Will write to network.
    dec := gob.NewDecoder(&network) // Will read from network.

    p := P{3, 4, 5, "Pythagoras"}
    
    err := enc.Encode(p)
    if err != nil {
        log.Fatal("encode error:", err)
    }

    var q Q
    err = dec.Decode(&q)
    if err != nil {
        log.Fatal("decode error:", err)
    }

    fmt.Printf("%#v\n\n", p)
    fmt.Printf("%#v\n\n", enc)
    fmt.Printf("%#v\n\n", dec)
    fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
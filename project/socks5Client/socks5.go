package main

import (
	"fmt"
	"github.com/gamexg/proxyclient"
	"io"
	"io/ioutil"
)

func main() {
	p, err := proxyclient.NewProxyClient("socks5://180.119.140.132:3245")
	if err != nil {
		panic(err)
	}

	c, err := p.Dial("tcp", "www.163.com:80")
	if err != nil {
		panic(err)
	}

	io.WriteString(c, "GET / HTTP/1.0\r\nHOST:www.163.com\r\n\r\n")
	b, err := ioutil.ReadAll(c)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
}

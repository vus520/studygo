package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/vus520/studygo/utils"
)

var (
	IpApiUrl = "http://ip.taobao.com/service/getIpInfo2.php?ip=8.8.8.8"
)

func test0(data string) {
	iter := jsoniter.ParseString(data)
	r := iter.Read()

	fmt.Println(r)
}

func test1(data string) {
	iter := jsoniter.ParseString(data)
	r := iter.Read()

	fmt.Println(r)
}

func main() {


	data, err := utils.FileGetContents(IpApiUrl)

	if err != nil {
		panic("FileGetContents returns " + err.Error())
		return
	}

	test0(data)
	test1(data)
}

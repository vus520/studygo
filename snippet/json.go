package main

import (
	"fmt"
	"github.com/json-iterator/go"
	"github.com/vus520/studygo/utils"
	"os"
)

var (
	IpApiUrl = "http://ip.taobao.com/service/getIpInfo2.php?ip=myip"
)

func main() {

	data, err := utils.FileGetContents(IpApiUrl)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	iter := jsoniter.ParseString(data)
	r := iter.Read()

	fmt.Println(r)

}

package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	IpApiUrl = "http://ip.taobao.com/service/getIpInfo2.php?ip=myip"
)

func main() {
	bodyByte := curl(IpApiUrl)
	dataJson := jsonDecode(bodyByte)
	fmt.Printf("%#v", dataJson["data"])

	ip := dataJson["data"]
	fmt.Printf("%#v", ip)
}

func curl(url string) []byte {

	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		fmt.Println("http.NewRequest error: ", err.Error())
		os.Exit(0)
	}

	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Add("Accept-Encoding", "gzip")
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Referer", "http://taobao.com/")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")

	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("http.Client.do error ", err.Error())
		os.Exit(1)
	}

	if response.StatusCode >= 400 {
		fmt.Println("http.code error: ", response.Status)
		os.Exit(1)
	}

	//需要在 switch 外面声明 bodyByte , switch 中声明的 bodyByte 为局部变量
	var bodyByte []byte
	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(response.Body)
		defer reader.Close()

		bodyByte, _ = ioutil.ReadAll(reader)
	default:
		bodyByte, _ = ioutil.ReadAll(response.Body)
	}

	return bodyByte
}

func jsonDecode(bodyByte []byte) map[string]interface{} {
	body := make(map[string]interface{})
	err := json.Unmarshal(bodyByte, &body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return nil
	}

	return body
}

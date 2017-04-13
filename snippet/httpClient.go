package main

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
)

var (
	Url    = "http://ip.taobao.com/service/getIpInfo.php?ip=myip"
	Header = map[string]string{
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
		"Connection":      "keep-alive",
		"Accept-Encoding": "gzip",
		"User-Agent":      "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0",
	}
)

type IpLocation struct {
	country,
	country_id,
	area,
	area_id,
	region,
	region_id,
	city,
	city_id,
	county_id,
	isp,
	isp_id,
	ip string
}

func main() {
	bodyByte, err := curl(Url, Header)

	if err != nil {
		fmt.Println("curl error:" + err.Error())
		os.Exit(1)
	}

	dataJson, err := jsonDecode(bodyByte)

	if err != nil {
		fmt.Println("jsonDecode error:" + err.Error())
		os.Exit(1)
	}

	data := dataJson["data"]
	fmt.Printf("%#v\n\n", data)

	dataMap := (data).(map[string]interface{})
	fmt.Printf("IpLocation: %s: %s%s\n\n", dataMap["ip"], dataMap["country"], dataMap["region"])

	for index, element := range data.(map[string]interface{}) {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d]\t\t\t\t,value is %d\n", index, value)
		default:
			fmt.Printf("list[%s]\t\t\t\t,value is %s\n", index, value)
		}
	}

	ip, _ := getValue(data, "ip")
	country, _ := getValue(data, "country")
	fmt.Printf("IP: %s, location: %s", ip, country)
}

func curl(url string, Header map[string]string) ([]byte, error) {

	client := &http.Client{}
	reqest, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		fmt.Println("http.NewRequest error: ", err.Error())
		os.Exit(0)
	}

	for k, v := range Header {
		reqest.Header.Add(k, v)
	}

	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("http.Client.do error ", err.Error())
	}

	if response.StatusCode >= 400 {
		return nil, errors.New("http.StatusCode: " + response.Status)
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

	return bodyByte, nil
}

func jsonDecode(bodyByte []byte) (map[string]interface{}, error) {
	body := make(map[string]interface{})
	err := json.Unmarshal(bodyByte, &body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getValue(d interface{}, label string) (interface{}, bool) {

    switch reflect.TypeOf(d).Kind() {
    case reflect.Struct:
        v := reflect.ValueOf(d).FieldByName(label)
             return v.Interface(), true

    case reflect.Map:
		inter := reflect.ValueOf(d).Interface()
		dict  := inter.(map[string]interface{})

		if val, ok := dict[label]; ok {
		    return val, ok
		}

        return nil, false
    }
   return nil, false
}

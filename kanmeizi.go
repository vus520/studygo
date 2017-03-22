/*

抓点妹子图

技术点
	goroutine
		WaitGroup
	http
		get
	regexp
		group
	file
		dir
		read, write
	variable
		loop
		printf
	import
		get
*/

package main

import (
	"crypto/md5"
	"fmt"
	"github.com/vus520/studygo/utils"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"sync"
	"time"
)

var imgList = make([]interface{}, 0, 1)
var pageList = make([]interface{}, 0, 1)
var timeStart = time.Now().Unix()
var w sync.WaitGroup

func main() {

	fmt.Println("生成任务，开始抓取")
	runtime.GOMAXPROCS(8)

	os.Mkdir("tmp", 0777)

	for i := 1; i < 2; i++ {
		//任务计数器增加
		w.Add(1)

		go func(i int) {
			url := fmt.Sprintf("http://www.kanmeizi.cn/tag_%d_1_16.html", i)

			fmt.Printf("Job: %s\n", url)

			body := curl(url)
			format(body)

			//任务计数器完成
			w.Done()
		}(i)
	}

	//等待任务计数器完成并清空，退出进程
	w.Wait()

	for i := range pageList {
		w.Add(1)

		go func(i int) {
			url := fmt.Sprintf("http://www.kanmeizi.cn%s", pageList[i])

			fmt.Printf("Job: %s\n", url)

			body := curl(url)
			format(body)

			//任务计数器完成
			w.Done()
		}(i)
	}

	w.Wait()

	fmt.Printf("页面抓取完成，获取图片: %d 张, 用时: %d 秒\n", len(imgList), time.Now().Unix()-timeStart)

	imgList = utils.Slice_unique(imgList)

	for i := range imgList {
		url := fmt.Sprintf("%s", imgList[i])
		img := curl(url)

		file := fmt.Sprintf("./tmp/%x.png", md5.Sum([]byte(img)))

		utils.FilePutContents(string(file), img)
	}

	fmt.Printf("下载图片: %d, 用时: %d s\n", len(imgList), time.Now().Unix()-timeStart)
}

//格式化页面，读取图片地址，分图片地址，存入全局变量
func format(body string) {
	r, _ := regexp.Compile(`<img class="height_min"[^>]+src="(?P<src>.*?)"`)

	// Compile vs MustCompile
	// FindAllStringSubmatch vs FindAllString vs FindStringSubmatch
	img := r.FindAllStringSubmatch(body, -1)

	for i := range img {
		imgList = append(imgList, img[i][1])
	}

	r, _ = regexp.Compile(`<a href="([^"]+)" data-page="\d+">`)

	// Compile vs MustCompile
	// FindAllStringSubmatch vs FindAllString vs FindStringSubmatch
	page := r.FindAllStringSubmatch(body, -1)

	for i := range page {
		pageList = append(pageList, page[i][1])
	}
}

func curl(url string) string {

	resp, err := http.Get(url)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Get faile: " + url)
		return ""
	}

	return string(body)
}

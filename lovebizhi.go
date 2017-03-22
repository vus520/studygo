package main

import (
    "fmt"
    "github.com/bitly/go-simplejson"
    "github.com/vus520/studygo/utils"
    "os"
    "path/filepath"
    "runtime"
    "strconv"
    "strings"
    "sync"
    "time"
)

var (
    DataRoot = "./tmp/lovebizhi/"
    PageUrl  = "http://api.lovebizhi.com/macos_v4.php?a=category&spdy=1&tid=3&order=hot&device=105&uuid=436e4ddc389027ba3aef863a27f6e6f9&mode=0&retina=0&client_id=1008&device_id=31547324&model_id=105&size_id=0&channel_id=70001&screen_width=1920&screen_height=1200&bizhi_width=1920&bizhi_height=1200&version_code=19&language=zh-Hans&jailbreak=0&mac=&p={pid}"
    w        sync.WaitGroup
)

// 壁纸类型，有编号，长宽和URL
type Wallpaper struct {
    Pid    int
    Url    string
    Width  int
    Height int
}

// 将图片下载并保存到本地
func SaveImage(paper *Wallpaper) {
    //按分辨率目录保存图片
    Dirname := DataRoot + strconv.Itoa(paper.Width) + "x" + strconv.Itoa(paper.Height) + "/"
    if !utils.IsDir(Dirname) {
        os.MkdirAll(Dirname, 0755)
    }

    //根据URL文件名创建文件
    filename := Dirname + filepath.Base(paper.Url)
    if utils.IsFile(filename) {
        return
    }

    w.Add(1)
    timeStart := time.Now().Unix()

    Body, err := utils.FileGetContents(paper.Url)

    if err == nil {
        utils.FilePutContents(filename, Body)
    }

    timeEnd := time.Now().Unix()
    fmt.Printf("%d: %s, 用时: %d (%d-%d) 秒, %s\n", paper.Pid, paper.Url, timeEnd-timeStart, timeStart, timeEnd, err)

    w.Done()
}

func main() {
    runtime.GOMAXPROCS(12)

    for i := 1; i < 10; i++ {

        url := strings.Replace(PageUrl, "{pid}", strconv.Itoa(i), -1)
        fmt.Printf("Page %d: %s\n", i, url)

        body, err := utils.FileGetContents(url)

        if err != nil {
            fmt.Println(err)
            continue
        }

        js, err := simplejson.NewJson([]byte(body))

        //遍历data下的所有数据
        data := js.Get("data").MustArray()
        for _, v := range data {
            v := v.(map[string]interface{})
            for kk, vv := range v {
                if kk == "file_id" {
                    //这里 vv 是一个[]interface{} json.Number，不知道怎么取出值，这里用了比较傻的Sprintf
                    vv := fmt.Sprintf("%s", vv)
                    imgid, _ := strconv.Atoi(vv)
                    url := fmt.Sprintf("http://s.qdcdn.com/c/%d,1920,1200.jpg", imgid)

                    paper := &Wallpaper{imgid, url, 1920, 1200}
                    go SaveImage(paper)
                }
            }
        }

    }

    w.Wait()
    fmt.Println("oh yes, all job done.")
}

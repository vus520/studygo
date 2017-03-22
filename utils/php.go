package utils

import (
	"io/ioutil"
	"net/http"
	"os"
)

//检查文件是否存在
func IsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

//检查目录是否存在
func IsDir(path string) bool {
	p, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return p.IsDir()
	}
}

//写内容到文件
func FilePutContents(filename string, content string) (val bool, err error) {
	if IsFile(filename) {
		return
	}

	fout, err := os.Create(string(filename))
	defer fout.Close()

	if err != nil {
		return false, err
	}
	fout.WriteString(content)

	return true, nil
}

func FileGetContents(url string) (str string, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

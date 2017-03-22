package utils

import (
	"os"
)

//检查文件是否存在
func IsFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

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

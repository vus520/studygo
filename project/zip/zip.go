package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const MODE_ZIP string = "zip"     //只备份到目标文件夹中对应的
const MODE_UNZIP string = "unzip" //只将源文件夹拷贝到对应目录，不进行备份   //TODO 不支持仅拷贝，太危险了

//传递两个参数进来
func main() {
	handleArgs()
}

func handleArgs() {
	args := os.Args
	length := len(args)
	if length < 4 {
		log.Println("args not enough")
		return
	}

	mode := os.Args[1]

	if !modeMatchs(mode) {
		panic("oh, bad mode !")
	}

	if strings.EqualFold(mode, MODE_ZIP) {
		sourceDirs := args[2 : length-1]
		//默认在当前目录创建压缩包
		targetZipFile := args[length-1]
		doZip(sourceDirs, targetZipFile)
	} else if strings.EqualFold(mode, MODE_UNZIP) {
		sourceZipFile := args[2]
		targetUnzipDir := args[3]
		doUnzip(sourceZipFile, targetUnzipDir)
	}
}

func doUnzip(sourceZipFile, targetUnzipDir string) {
	if !strings.HasSuffix(sourceZipFile, ".zip") {
		panic("unzip source file must be zip format")
	}

	if err := makeParent(targetUnzipDir); err != nil {
		fmt.Printf("make target parnet err: %s\n", err)
		panic("make target parent dir failed")
	}

	log.Println("doUnzip --- sourceZipFile: " + sourceZipFile)
	log.Println("doUnzip --- targetUnzipDir: " + targetUnzipDir)
	DeCompress(sourceZipFile, targetUnzipDir)
}

func doZip(sourceDirs []string, targetZipFile string) {
	log.Println("targetZipFile: " + targetZipFile)

	if !strings.HasSuffix(targetZipFile, ".zip") {
		panic("target file must be zip format")
	}

	if err := makeParent(targetZipFile); err != nil {
		fmt.Printf("make parnet err: %s\n", err)
		panic("make target parent dir failed")
	}

	var files []*os.File
	for _, f := range sourceDirs {
		if dir, err := os.Open(f); err != nil {
			fmt.Printf("sourceFile open failed: %s\n", err)
			panic("sourceFile open failed")
		} else {
			files = append(files, dir)
		}
	}

	Compress(files, targetZipFile)
}

func makeParent(dir string) error {
	parent := filepath.Dir(dir)
	if exists, _ := PathExists(parent); exists {
		return nil
	}
	fmt.Println("make parent: " + parent)

	return makeSureDir(parent)
}

func makeSureDir(dir string) error {
	if exists, _ := PathExists(dir); !exists {
		// err := os.Mkdir(bkDir, os.ModePerm)
		fmt.Println("makeing dir: " + dir)
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			return err
		}
	}
	return nil
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//压缩文件
//files 文件数组，可以是不同dir下的文件或者文件夹
//dest 压缩文件存放地址
func Compress(files []*os.File, dest string) error {
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

//解压
func DeCompress(zipFile, dest string) error {
	dest = dest + "/";
	dest = strings.Replace(dest, "\\", "/", -1)
	dest = strings.Replace(dest, "//", "/", -1)
	log.Printf("unzip %s to %s", zipFile, dest)
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()

	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
	
		filename := dest + file.Name
		log.Println(filename)
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}

		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()

		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}

		w.Close()
		rc.Close()
	}
	
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}

//测试阶段，只支持bk和bkc
func modeMatchs(mode string) bool {
	// if strings.EqualFold(mode, MODE_BK) || strings.EqualFold(mode, MODE_BKCOPY) || strings.EqualFold(mode, MODE_COPY) {
	if strings.EqualFold(mode, MODE_ZIP) || strings.EqualFold(mode, MODE_UNZIP) {
		return true
	}

	return false
}

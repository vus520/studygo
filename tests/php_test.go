package tests

import (
	"fmt"
	"github.com/vus520/studygo/utils"
	"runtime"
	"testing"
	"time"
)

func Test_FileGetContents(t *testing.T) {

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {

		if !utils.IsFile("/etc/passwd") {
			t.Error("/etc/password not found on Linux system")
			t.Fail()
		}

		if utils.IsFile("/youmusthavethisfileondisk-_-") {
			t.Error("file exists is NOT-reasonable.")
			t.Fail()
		}

	} else {
		t.Skip("not support system")
	}

}

func Test_FilePutContents(t *testing.T) {

	tmpfile := "/tmp/gotestputcontents.tmp"

	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		timestamp := fmt.Sprintf("randstr: %d", time.Now().Unix())
		utils.FilePutContents(tmpfile, timestamp)

		data, _ := utils.FileGetContents(tmpfile)
		utils.Unlink(tmpfile)

		if timestamp != data {
			t.Error("file put contents not match read contents.")
			t.Fail()
		}
	} else {
		t.Skip("not support system")
	}

}

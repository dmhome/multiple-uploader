package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"github.com/axgle/mahonia"
	"os"
	"io/ioutil"
)

func GetLogId() string {
	time.Sleep(1*time.Microsecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rint := r.Intn(10000)
	fint := fmt.Sprintf("%05d", rint)
	return strconv.FormatInt(time.Now().UnixNano(), 10) + fint
}

func GetcurDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func ConvertToString(src string, srcCode string, tagCode string) []byte {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func FileReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
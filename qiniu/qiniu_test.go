package qiniu

import (
	"fmt"
	"ksong/utils"
	"testing"
)

func TestUpload(t *testing.T) {
	filePath := "../music/original/五月天-伤心的人别听慢歌.txt"
	fileKey := utils.GetLogId() + ".txt"
	flag := Upload(filePath, fileKey)
	fmt.Println(flag)
	fmt.Println(fileKey)
}

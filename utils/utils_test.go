package utils

import (
	"fmt"
	"testing"
)

func TestGetLogId(t *testing.T) {
	logId := GetLogId()
	fmt.Println(logId)
}

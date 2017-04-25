package vanConfig

import (
	"testing"
)

func Test_Config(t *testing.T) {

	configJson := Config{FilePath:"E:/go/2.json",FileType:"json"}

	err := configJson.Open()

	if err != nil{
		t.Error("不能通过") // 如果不是如预期的那么就报错
	}
}
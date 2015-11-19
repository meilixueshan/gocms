package utils

import (
	"bytes"
	"strings"
)

//模板格式化函数
func CategoryFormat(name string, innerCode string) string {
	arr := strings.Split(innerCode, ".")
	var buf bytes.Buffer
	for i := 1; i < len(arr); i++ {
		buf.WriteString("───")
	}
	buf.WriteString(name)
	return buf.String()
}
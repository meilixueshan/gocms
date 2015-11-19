package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"strings"
	"time"
	"strconv"
)


//加密成md5的字符串
func ToMD5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}


//对象转json字符串
func ToJson(v interface{}) string {
	str ,err := json.Marshal(&v)
	if err != nil {
		return "{Success:false, Msg:\"\", Error:\"" + err.Error() + "\"}"
	}
	return string(str);
}

//时间格式化（如：yyyy-mm-dd hh:ii:ss或y-m-d h:i:s）
func TimeFormat(tm time.Time, format string) string {
	format = strings.ToLower(format)	//全部转为小写
	
	format = strings.Replace(format, "yyyy", "2006", -1)
	format = strings.Replace(format, "y", "06", -1)
	
	format = strings.Replace(format, "hh", "15", -1)
	format = strings.Replace(format, "h", "3", -1)
	
	format = strings.Replace(format, "mm", "01", -1)
	format = strings.Replace(format, "m", "1", -1)
	
	format = strings.Replace(format, "dd", "02", -1)
	format = strings.Replace(format, "d", "2", -1)
	
	format = strings.Replace(format, "ii", "04", -1)
	format = strings.Replace(format, "i", "4", -1)
	
	format = strings.Replace(format, "ss", "05", -1)
	format = strings.Replace(format, "s", "5", -1)
	
	//go语言的时间格式化很特别，采用的是"2006-01-02 15:04:05"，可以使用简单的方式来记，就是1，2，3，4，5，6
	//外国的日期时间格式顺序一般是：月，日，时，分，秒，年，对应的就是1,2,3,4,5,6
	return tm.Format(format)
}

func Concat(args ...string) string {
	var buf bytes.Buffer
	for _, n := range args {
		buf.WriteString(n)
	}
	return buf.String()
}

func IsNullOrEmpty(str string) bool {
	if strings.EqualFold(str, "") {
		return true
	}
	return false
}

//数字格式化（如：NumberFormat(1, "000")，那么结果就是001）
func NumberFormat(value int, format string) string {
	val := strconv.Itoa(value)
    formatLen := len(format)
    valueLen := len(val)
	var buf bytes.Buffer
    if valueLen < formatLen {
		for i := 0; i < formatLen - valueLen; i++ {
			buf.WriteString("0")
        }
    }
	
	buf.WriteString(val)
    return buf.String()
}
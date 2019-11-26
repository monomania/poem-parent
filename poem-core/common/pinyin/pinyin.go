package pinyin

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
)

var pyValue_length = len(pyValue)
var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

// transform UTF-8 string to GBK string and replace it, if transformed success, returned nil error, or died by error message
func Utf8ToGbk(str *string) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader([]byte(*str)), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	*str = string(b)
	return
}

//get chinese pinyin number code
//param s must be chinese character with utf8 encoding
func Code(s string) int {
	gbkString, _ := Utf8ToGbk(&s)
	var i1, i2 int
	i1 = int(gbkString[0])
	i2 = int(gbkString[1])
	return i1*256 + i2 - 65536
}

// convert chinese to pinyin
func Convert(s string) string {
	pyString := ""
	var str string
	var code int

	for _, rune := range s {
		str = string(rune)
		if hzRegexp.MatchString(str) { //chinese
			code = Code(str)
			if code > 0 && code < 160 {
				pyString += str
			} else {
				if v, ok := tableMap[code]; ok { //map by table
					pyString += v
				} else {

					for i := (pyValue_length - 1); i >= 0; i-- {

						if pyValue[i] <= code {
							pyString += pyName[i]
							break
						}
					}
				}
			}
		} else { //other
			pyString += str
		}
	}

	return pyString
}

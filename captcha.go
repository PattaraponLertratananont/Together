package captcha

import (
	"strconv"
)

func ReturnRight(right int) string {
	StrNumber := []string{" + Zero", " + One", " + Two", " + Three", " + Four", " + Five", " + Six", " + Seven", " + Eight", " + Nine"}
	return StrNumber[right]
}
func Captcha(pattern, left, operator, right int) string {
	return strconv.Itoa(left) + ReturnRight(right)
}

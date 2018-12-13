package captcha

import (
	"strconv"
)

func Captcha(pattern, left, operator, right int) string {
	return strconv.Itoa(left) + " + One"
}

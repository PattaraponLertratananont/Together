package captcha

import (
	"strconv"
)

func Captcha(pattern, left, operator, right int) string {
	if right == 0 {
		return strconv.Itoa(left) + " + Zero"
	}
	return strconv.Itoa(left) + " + One"
}

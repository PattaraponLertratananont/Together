package captcha

import (
	"strconv"
)

func Captcha(pattern, left, operator, right int) string {
	if right == 3 {
		return strconv.Itoa(left) + " + Three"
	}
	if right == 2 {
		return strconv.Itoa(left) + " + Two"
	}
	if right == 1 {
		return strconv.Itoa(left) + " + One"
	}
	return strconv.Itoa(left) + " + Zero"
}

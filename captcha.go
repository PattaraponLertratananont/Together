package captcha

import (
	"strconv"
)

func Captcha(pattern, left, operator, right int) string {
	if right == 9 {
		return strconv.Itoa(left) + " + Nine"
	}
	if right == 8 {
		return strconv.Itoa(left) + " + Eight"
	}
	if right == 7 {
		return strconv.Itoa(left) + " + Seven"
	}
	if right == 6 {
		return strconv.Itoa(left) + " + Six"
	}
	if right == 5 {
		return strconv.Itoa(left) + " + Five"
	}
	if right == 4 {
		return strconv.Itoa(left) + " + Four"
	}
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

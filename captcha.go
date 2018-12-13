package captcha

import (
	"strconv"
)

func Operators(operator int) string {
	StrOperator := []string{" + ", " - ", " * "}
	return StrOperator[operator-1]
}
func ReturnStrNumber(number int) string {
	StrNumber := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return StrNumber[number]
}
func Captcha(pattern, left, operator, right int) string {
	if pattern == 1 {
		return strconv.Itoa(left) + Operators(operator) + ReturnStrNumber(right)
	} else if pattern == 2 {
		return ReturnStrNumber(left) + Operators(operator) + strconv.Itoa(right)
	} else {
		return "Wrong Input !!! Please Try Again..."
	}
}

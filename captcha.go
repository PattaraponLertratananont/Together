package captcha

import (
	"strconv"
)

func Operators(operator int) string {
	StrOperator := []string{" + ", " - ", " * "}
	return StrOperator[operator-1]
}
func ReturnRight(right int) string {
	StrNumber := []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return StrNumber[right]
}
func Captcha(pattern, left, operator, right int) string {
	return strconv.Itoa(left) + Operators(operator) + ReturnRight(right)
}

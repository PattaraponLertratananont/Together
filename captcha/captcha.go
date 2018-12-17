package main

import (
	"fmt"
	"strconv"
)

//Try to push to muyonz
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

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)
	fmt.Println("Output = " + Captcha(a, b, c, d))
}

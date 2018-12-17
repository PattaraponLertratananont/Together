package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Please Input Number1 : ")
	fmt.Scan(&num1)
	fmt.Print("Please Input Number2 : ")
	fmt.Scan(&num2)
	fmt.Println(num1, "+", num2, "=", num1+num2)
	fmt.Println(num1, "-", num2, "=", num1-num2)
	fmt.Println(num1, "*", num2, "=", num1*num2)
	fmt.Println(num1, "/", num2, "=", num1/num2)
	fmt.Println(num1, "%", num2, "=", num1%num2)
}

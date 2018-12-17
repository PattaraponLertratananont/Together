package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Intput Number1 Number2 (Ex 7 3) : ")
	fmt.Scan(&num1, &num2)
	Cal(num1, num2)
}
func Cal(num1, num2 int) {
	Plus := num1 + num2
	fmt.Printf("%v + %v = %v\n", num1, num2, Plus)
	Minut := num1 - num2
	fmt.Printf("%v - %v = %v\n", num1, num2, Minut)
	Divide := num1 / num2
	fmt.Printf("%v / %v = %v\n", num1, num2, Divide)
	Multiple := num1 * num2
	fmt.Printf("%v * %v = %v\n", num1, num2, Multiple)
	Mod := num1 % num2
	fmt.Printf("%v %% %v = %v\n", num1, num2, Mod)

}

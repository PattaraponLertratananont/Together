package main

//เขียนโปรแกรมinputเลข 2 ตัวแล้วแสดงค่าที่มากกว่าทางจอ
import (
	"fmt"
)

func main() {
	var first, second int

	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)

	if first < second {
		fmt.Println("Max: ", second)
	} else if first > second {
		fmt.Println("Max: ", first)
	} else {
		fmt.Println("It's equal.")
	}
}
func init() {
	fmt.Println("Hi, chapter4")
	fmt.Println("-------------------------------")
}

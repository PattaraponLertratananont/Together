package main

//เขียนโปรแกรมรับinputเลข 2ตัว แล้วแสดงผล หรม.ของเลขทั้งสอง
import "fmt"

func main() {
	var first, second int
	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)
}
func init() {
	fmt.Println("Hi, chapter15")
	fmt.Println("---------------------")
}

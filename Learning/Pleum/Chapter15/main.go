package main

//เขียนโปรแกรมรับinputเลข 2ตัว แล้วแสดงผล หรม.ของเลขทั้งสอง
import "fmt"

func main() {
	var first int
	p := 2
	sum := 1
	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	// fmt.Print("Input second number: ")
	// fmt.Scanln(&second)

	for first > 1 && p < first {
		if first%p == 0 {
			sum = sum * p
			first = first / p
			p = 2
		} else {
			p++
		}
	}
	fmt.Println(sum)
}
func init() {
	fmt.Println("Hi, chapter15")
	fmt.Println("---------------------")
}

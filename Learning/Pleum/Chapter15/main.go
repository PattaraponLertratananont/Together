package main

//เขียนโปรแกรมรับinputเลข 2ตัว แล้วแสดงผล หรม.ของเลขทั้งสอง
import "fmt"

func main() {
	var first, second int
	p := 2
	sum := 1
	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)

	for first > 1 && p < first && second > 1 && p < second {
		if first%p == 0 && second%p == 0 {
			sum = sum * p
			first = first / p
			second = second / p
			p = 2
		} else {
			p++
		}
	}
	fmt.Println("Greatest common divisor: ", sum)
}
func init() {
	fmt.Println("Hi, chapter15")
	fmt.Println("---------------------")
}

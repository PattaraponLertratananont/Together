package main

//เขียนโปรแกรมรับinputเลข 2ตัว แล้วแสดงผล หรม.ของเลขทั้งสอง
import "fmt"

func main() {
	var first, second int
	maxcd, gcd := 0, 0
	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)

	for maxcd = 1; maxcd < first && maxcd < second; maxcd++ {
		if first%maxcd == 0 && second%maxcd == 0 {
			gcd = maxcd
		}
	}
	fmt.Println("Greatest common divisor: ", gcd)
}

func init() {
	fmt.Println("Hi, chapter15")
	fmt.Println("---------------------")
}

package main

//เขียนโปรแกรมรับinputเลข 2ตัว แล้วแสดงผล หรม.ของเลขทั้งสอง
import "fmt"

func main() {
	var first, second int
	cd, gcd := 0, 0
	fmt.Print("Input first number: ")
	fmt.Scanln(&first)
	fmt.Print("Input second number: ")
	fmt.Scanln(&second)

	for cd = 1; cd < first && cd < second; cd++ {
		if first%cd == 0 && second%cd == 0 {
			gcd = cd
		}
	}
	fmt.Println("Greatest common divisor: ", gcd)
}

func init() {
	fmt.Println("Hi, chapter15")
	fmt.Println("---------------------")
}

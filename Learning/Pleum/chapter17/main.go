package main

//เขียนโปรแกรมรับเลขบวกinput 1ตัว แล้วแสดงsquare root หากใส่ติดลบให้ป้อนใหม่จนกว่าจะใส่บวก
import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Input 1 number: ")
	fmt.Scanln(&input)

	fmt.Printf("Square root %d = %.2f \n", int32(input), math.Sqrt(input))
}
func init() {
	fmt.Println("Hi, chapter 17 ")
	fmt.Println("------------------------")
}

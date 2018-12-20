package main

//เขียนโปรแกรมinputเลข 3ตัว แสดงผลว่าเป็นสามเหลี่ยมมุมฉากหรือไม่?
import (
	"fmt"
	"math"
)

func main() {
	var x, y, z float64

	fmt.Print("X: ")
	fmt.Scanln(&x)
	fmt.Print("Y: ")
	fmt.Scanln(&y)
	fmt.Print("Z: ")
	fmt.Scanln(&z)

	if math.Pow(x, 2)+math.Pow(y, 2) == math.Pow(z, 2) {
		fmt.Println("X, Y, Z are side of right triangle.")
	} else {
		fmt.Println("No, it's not side of right triangle. ")
	}
}
func init() {
	fmt.Println("Hi, chapter11")
	fmt.Println("--------------------------------")
}

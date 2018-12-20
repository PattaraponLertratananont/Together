package main

//เขียนโปรแกรมinputเลข 3ตัว แสดงผลว่าเป็นสามเหลี่ยมมุมฉากหรือไม่?
import "fmt"

func main() {
	var x, y, z int

	fmt.Print("X: ")
	fmt.Scanln(&x)
	fmt.Print("Y: ")
	fmt.Scanln(&y)
	fmt.Print("Z: ")
	fmt.Scanln(&z)
}
func init() {
	fmt.Println("Hi, chapter11")
	fmt.Println("--------------------------------")
}

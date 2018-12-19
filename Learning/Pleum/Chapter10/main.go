package main

//เขียนโปรแกรมinputเลข 3ตัวแล้วแสดงผลว่าเลขทั้งสามเป็นด้านของสามเหลี่ยมเหรือไม่
import "fmt"

func main() {
	var a, b, c int

	fmt.Print("A: ")
	fmt.Scanln(&a)
	fmt.Print("B: ")
	fmt.Scanln(&b)
	fmt.Print("C: ")
	fmt.Scanln(&c)
}
func init() {
	fmt.Println("Hi, chapter10")
	fmt.Println("--------------------------")
}

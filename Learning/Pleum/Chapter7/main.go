package main

//เขียนโปรแกรมinput 10ตัว แล้วแล้วแสดงผมรวมทั้งหมดและแสดงค่าเฉลี่ย
import "fmt"

func main() {
	var input float32

	for i := 1; i <= 10; i++ {
		fmt.Printf("Input number %d: ", i)
		fmt.Scanln(&input)
	}
}
func init() {
	fmt.Println("Hi, chapter7")
	fmt.Println("--------------------")
}

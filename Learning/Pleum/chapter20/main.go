package main

import "fmt"

//เขียนโปรแกรมinputเลข 1ตัว แล้วแสดงดอกจัน(*)เป็นสี่เหลี่ยม nxn แบบตารางหมากรุก
func main() {
	var keyboard int
	fmt.Print("Input number: ")
	fmt.Scanln(&keyboard)

	for i := 0; i < keyboard; i++ {
		for j := 0; j < keyboard; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

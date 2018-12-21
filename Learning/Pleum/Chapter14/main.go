package main

//เขียนโปรแกรมรับinputเลข 1ตัวแ้วแสดงผลการแยกตัวประกอบ
import "fmt"

func main() {
	input, p := 0, 2
	fmt.Print("Input number: ")
	fmt.Scanln(&input)
	fmt.Printf("factor %d= ", input)

	for input > 1 && p < input {
		if input%p == 0 {
			fmt.Printf("%d x ", p)
			input = input / p
			p = 2
		} else {
			p++
		}
	}
	fmt.Printf("%d \n", p)

}
func init() {
	fmt.Println("Hi, chapter14")
	fmt.Println("--------------------------------")
}

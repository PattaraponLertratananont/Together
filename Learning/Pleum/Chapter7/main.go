package main

//เขียนโปรแกรมinput 10ตัว แล้วแล้วแสดงผมรวมทั้งหมดและแสดงค่าเฉลี่ย
import "fmt"

func main() {
	var input, sum, average float64

	for i := 1; i <= 10; i++ {
		fmt.Printf("Input number %d: ", i)
		fmt.Scanln(&input)
		sum += input
	}
	average = sum / 10
	fmt.Println("Sum: ", sum)
	fmt.Println("Average: ", average)
}
func init() {
	fmt.Println("Hi, chapter7")
	fmt.Println("--------------------")
}

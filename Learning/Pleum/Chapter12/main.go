package main

//เขียนโปรแกรมรับinputเลข 10ตัวแล้วแสดงเลขที่อยู่ระหว่างเลขคู่
import "fmt"

func main() {
	box := [10]int{}

	for i := 0; i < len(box); i++ {
		fmt.Printf("Input number %d: ", i+1)
		fmt.Scanln(&box[i])
	}

}
func init() {
	fmt.Println("Hi, chapter 12")
	fmt.Println("-------------------------")
}

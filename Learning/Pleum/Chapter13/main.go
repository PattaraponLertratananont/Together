package main

//เขียนโปรแกรมรับinputเลข 10ตัวแล้วแสดงผลเรียงลำดับจากน้อย-มาก
import "fmt"

func main() {
	box := [10]int{}
	for i := 0; i < len(box); i++ {
		fmt.Printf("Input number %d: ", i)
		fmt.Scanln(&box[i])
	}
}
func init() {
	fmt.Println("Hi, chapter13")
	fmt.Println("---------------------------")
}

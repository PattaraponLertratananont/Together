package main

//เขียนโปรแกรมรับinputเลขจำนวนเต็มไปเรื่อยๆจนกว่าจะป้อน0 เมื่อรับ0แล้วโปรแกรมจะprintจำนวนตัวเลข
//พร้อมผลรวมแยกเป็นเลขคู่(even)และเลขคี่(odd)
import "fmt"

func main() {
	var input int
	for {
		fmt.Print("Input number: ")
		fmt.Scanln(&input)
		if input == 0 {

			break
		}
	}
}
func init() {
	fmt.Println("Hi, chapter8")
	fmt.Println("----------------------------")
}
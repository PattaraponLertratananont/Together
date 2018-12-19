package main

//เขียนโปรแกรมรับinputเลขจำนวนเต็มไปเรื่อยๆจนกว่าจะป้อน0 เมื่อรับ0แล้วโปรแกรมจะprintจำนวนตัวเลข
//พร้อมผลรวมแยกเป็นเลขคู่(even)และเลขคี่(odd)
import "fmt"

func main() {
	var input, even, odd int
	sumeven, sumodd := 0, 0
	for {
		fmt.Print("Input number: ")
		fmt.Scanln(&input)
		if input%2 == 0 {
			sumeven += input
			even++
		} else {
			sumodd += input
			odd++
		}
		if input == 0 {
			fmt.Printf("Sum %d even numbers: %d\n", even-1, sumeven)
			fmt.Printf("Sum %d odd numbers: %d\n", odd, sumodd)
			break
		}
	}
}
func init() {
	fmt.Println("Hi, chapter8")
	fmt.Println("----------------------------")
}

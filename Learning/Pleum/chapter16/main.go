package main

//เขียนโปรแกรมรับinputเลข 1ตัว แสดงผลว่าเป็นจำนวนเฉพาะ(Prime number)หรือไม่
import "fmt"

func main() {
	var input, count int
	fmt.Print("Input number: ")
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		if input%i == 0 {
			count++
		}
	}

}

func init() {
	fmt.Println("Hi, chapter16")
	fmt.Println("--------------------------")
}

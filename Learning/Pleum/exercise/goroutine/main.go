package main

import (
	"fmt"
	"time"
)

func buyGkassesMinimart() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่นตา ที่มินิมาท")
}
func buyWatchesCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา ที่เซนทรัล")
}
func buyFruitsParagon() {
	time.Sleep(1 * time.Second)
}
func buyCarHonda() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ ที่ฮอนด้า")
}

func main() {
	buyGkassesMinimart()
	buyWatchesCentral()
	buyFruitsParagon()
	buyCarHonda()
}
func init() {
	fmt.Println("Hi, goroutine")
	fmt.Println("-----------------------------")
}

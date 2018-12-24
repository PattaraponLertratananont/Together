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
	fmt.Println("ซื้อผลไม้ ที่พารากอน")
}
func buyCarHonda() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ ที่ฮอนด้า")
}

func main() {
	start := time.Now()
	go buyGkassesMinimart()
	go buyWatchesCentral()
	buyFruitsParagon()
	buyCarHonda()
	fmt.Println("total time: ", time.Since(start))
}
func init() {
	fmt.Println("Hi, goroutine")
	fmt.Println("-----------------------------")
}

package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("TeamA: crawl")
	fmt.Println("TeamB: run")
	fmt.Println("TeamC: fly")
	time.Sleep(1 * time.Second)
}
func init() {
	fmt.Println("Hi, goroutine")
	fmt.Println("-----------------------------")
}

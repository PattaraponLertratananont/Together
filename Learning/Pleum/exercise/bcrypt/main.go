package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	secret := `123456`
	byteslice, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(secret)
	fmt.Println(byteslice)

}
func init() {
	fmt.Println("Hi, bcrypt golang.")
	fmt.Println("------------------------------")
}

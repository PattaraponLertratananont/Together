package main

import "fmt"

func main() {
	fname, lname, gender := "", "", ""
	age := 0
	grade := 0.0

	fmt.Print("Please Input Your FirstName And LastName  : ")
	fmt.Scan(&fname)
	fmt.Scan(&lname)
	fmt.Print("Please Input Your Ages : ")
	fmt.Scan(&age)
	fmt.Print("Please Input Your Gender (M/Fm) : ")
	fmt.Scan(&gender)
	fmt.Print("Please Input Your Grade : ")
	fmt.Scan(&grade)

	fmt.Println("Name : "+fname+" "+lname+"\n"+"Age : ", age, "\n"+"Gender : "+gender+"\n"+"Grade : ", grade)

}

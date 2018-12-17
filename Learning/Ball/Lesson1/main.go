package main

import "fmt"

func main() {
	Profile()
}

func Profile() {
	Name, Subname, Gender := "", "", ""
	fmt.Print("Name : ")
	fmt.Scan(&Name, &Subname)
	fmt.Print("Gender : ")
	fmt.Scan(&Gender)
	fmt.Print("Age : ")
	var Age int
	fmt.Scan(&Age)
	fmt.Print("Grade : ")
	var Grade float32
	fmt.Scan(&Grade)
	fmt.Printf("-----------------------\nName : %v %v\nGender : %v\nAge : %v\nGrade : %v\n-----------------------",
		Name, Subname, Gender, Age, Grade)
}

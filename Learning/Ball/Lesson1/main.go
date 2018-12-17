package main

import "fmt"

func main() {
	Profile()
}

func Profile() {
	fmt.Print("Name : ")
	Name := ""
	fmt.Scan(&Name)
	fmt.Print("Sub name : ")
	Subname := ""
	fmt.Scan(&Subname)
	fmt.Print("Gender : ")
	Gender := ""
	fmt.Scan(&Gender)
	fmt.Print("Age : ")
	var Age int
	fmt.Scan(&Age)
	fmt.Print("Grade : ")
	var Grade float32
	fmt.Scan(&Grade)

	fmt.Printf("-----------------------\nName : %v %v\nGender : %v\nAge : %v\nGrade : %v\n-----------------------", Name, Subname, Gender, Age, Grade)
}

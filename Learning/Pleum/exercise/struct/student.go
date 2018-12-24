package main

import "fmt"

type student struct {
	id        string
	firstname string
	lastname  string
	no        int
	phone     string
}

func main() {
	studentLists := []student{}

	student1 := student{
		id:        "605233000",
		firstname: "Tammarut",
		lastname:  "Nualmeechue",
		no:        5,
		phone:     "087-1542740",
	}
	student2 := student{
		id:        "5099992322",
		firstname: "Bruce",
		lastname:  "Lee",
		no:        6,
		phone:     "094-554618",
	}
	student3 := student{
		id:        "1011101010",
		firstname: "The",
		lastname:  "Flash",
		no:        99,
		phone:     "011-1111111",
	}
	fmt.Println("Before: ", studentLists)
	studentLists = append(studentLists, student1)
	studentLists = append(studentLists, student2)
	studentLists = append(studentLists, student3)

	fmt.Println("After: ", studentLists)
}
func init() {
	fmt.Println("Hi, student")
	fmt.Println("--------------------------")
}

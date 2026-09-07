package main 

import "fmt"

func main() {
	person := map[string] string{
		"name": "Tegar",
		"age": "20",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	book := make(map[string] string)
	book["title"] = "Belajar Golang"
	book["author"] = "Tegar Alfa Rizzi"
	book["test"] = "cihuy"
	
	fmt.Println(book)

	delete(book, "test")
	fmt.Println(book)
}
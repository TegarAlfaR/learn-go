package main

import "fmt"

func main() {
	const firstName = "Tegar"
	const middleName = "Alfa"

	fmt.Println(firstName, middleName)

	const (
		name1 = "Tegar"
		name2 = "Alfa"
		name3 = "Rizzi"
	)

	fmt.Println(name1, name2, name3)
}
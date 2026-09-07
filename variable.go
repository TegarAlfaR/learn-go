package main

import "fmt"

func main() {
	var name string

	name  = "Tegar Alfa Rizzi"

	fmt.Println(name)
	
	name = "budi santoso"
	fmt.Println(name)

	var name2 = "mr cihuy"
	fmt.Println(name2)

	var number = 10
	fmt.Println(number)

	// pake := biar ngga perlu deklare var(hanya bisa di deklarasi awal)
	number2 := 22
	fmt.Println(number2)

	number2 = 222
	fmt.Println(number2)

	var (
		firstName = "Tegar"
		lastName = "Rizzi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}

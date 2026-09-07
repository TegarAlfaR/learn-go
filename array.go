package main 

import "fmt"

func main() {
	var names [3] string
	names[0] = "Tegar"
	names[1] = "Alfa"
	names[2] = "Rizzi"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names)

	var value = [3]int{
		123,
		456,
		789,
	}
	fmt.Println(value)
	fmt.Println(len(value))
	value[2] = 1000
	fmt.Println(value)
	
	// pake [...] harus di deklarasikan nilai/valuenya
	var number = [...]int{1,2,3,4,5,6}
	fmt.Println(number)
	fmt.Println(len(value))

 }
package main

import "fmt"

func fullName()(string, string){
	return "Tegar", "Alfa Rizzi"
}

func main(){
	fmt.Println(fullName())

	// jika hanya butuh 1 return saja
	firstName, _ := fullName()
	fmt.Println(firstName)
}
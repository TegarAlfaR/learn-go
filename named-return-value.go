package main

import "fmt"

func names()(firstName, middleName, lastName string){
	firstName = "Tegar"
	middleName = "Alfa"
	lastName = "Rizzi"
	return firstName, middleName, lastName
}

func main(){
	firstName, middleName, lastName := names()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
	fmt.Println(firstName, middleName, lastName)

}
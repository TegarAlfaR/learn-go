package main

import "fmt"

func login(){
	fmt.Println("selesai func login")
}

func application(){

	defer login()	

	fmt.Println("selesai func application")
}

func main() {
	application()
}
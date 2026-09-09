package main

import "fmt"

func endApp(){
	fmt.Println("end app")
	message := recover()
	fmt.Println("error message :", message)
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("ERROR CIHUY")
	}
}


func main(){
	runApp(true)
}
package main

import "fmt"

func endApp(){
	fmt.Println("end app")
}

func runApp(error bool){
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main(){
	runApp(true)
}
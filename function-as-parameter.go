package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter){
	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}
// func sayHelloWithFilter(name string, filter func(string) string){
// 	filteredName := filter(name)
// 	fmt.Println("hello", filteredName)
// }

func spamFilter(name string) string{
	if name == "kambing"{
		return "..."
	}else{
		return name
	}
}


func main(){
	sayHelloWithFilter("kuda", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("kambing", filter)
}
package main

import "fmt"

func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	}else{
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	data := NewMap("sadad")
	if data == nil{
		fmt.Println("data is empty")
	}else{
		fmt.Println(data)
		fmt.Println(data["name"])
	}
}
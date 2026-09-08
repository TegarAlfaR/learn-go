package main

import "fmt"

// import "fmt"

func sayHay(name string)string{
	hello := "hello " + name
	return hello
}

func main(){
	result :=sayHay("tegar")
	fmt.Println(result)
}
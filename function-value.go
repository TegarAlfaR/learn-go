package main

import "fmt"

func sayHello(name string) string {
	return "hello" + name
}

func main() {
	hello := sayHello
	result := hello("tegar")
	fmt.Println(result)

	fmt.Println(hello("budi"))
}
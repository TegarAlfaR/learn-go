package main

import "fmt"

func Test() any{
	// return 1
	return true
}

func main() {
	var coba any = Test()
	fmt.Println(coba)
}
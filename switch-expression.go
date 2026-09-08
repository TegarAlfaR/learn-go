package main 

import "fmt"

func main() {
	name := "tegar"

	switch name{
	case "tegar":
		fmt.Println("hello", name)
	case "budi":
		fmt.Println("hello budi")	
	default:
		fmt.Println("hello world")
	}

	switch length := len(name); length > 4 {
	case true:
		fmt.Println("terlalu panjang")
	case false:
		fmt.Println("sudah oke")
	}
}
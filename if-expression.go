package main 

import "fmt"

func main() {
	name := "tegar"

	if name == "tegar1" {
		fmt.Println("hello", name)
	}else if name == "tegar2" {
		fmt.Println("hello tegar2")	
	}else{
		fmt.Println("hello world")
	}

	if length := len(name); length > 5{
		fmt.Println("terlalu panjang")
	}else{
		fmt.Println("sudah oke")
	}
}
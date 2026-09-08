package main

import "fmt"

func main(){
	for i:= 1; i <= 10; i++{
		if i == 6{
			break;
		}
		fmt.Println("for ke -", i)
	}

	for i:= 1; i <= 10; i++{
		if i%2 == 0{
			continue;
		}

		fmt.Println("perulangan ke-", i)
	}
}
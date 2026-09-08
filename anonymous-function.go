package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList){
	if blackList(name){
		fmt.Println("you are blocked", name)
	}else{
		fmt.Println("welcome", name)
	}
}

func main() {

	// contoh anonymous func 1
	blacklist := func(name string) bool {
		return name == "anjay"
	}

	registerUser("tegar", blacklist)


	// contoh anonymous func 2
	registerUser("anjay", func(name string) bool{
		return name == "anjay"
	})

}
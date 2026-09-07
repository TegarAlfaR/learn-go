package main

import "fmt"

func main() {

	type noKtp string

	var ktpTegar noKtp = "1234567890"
	fmt.Println(ktpTegar)

	var contoh = "99109090909"
	fmt.Println(noKtp(contoh))

	
}
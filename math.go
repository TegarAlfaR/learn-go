package main

import "fmt"

func main() {

	var a = 10
	var b = 5
	var c = 2
	var d = a+b*c

	fmt.Println("hasil = ", d)

	a += 10
	fmt.Println("hasil = ", a)
	b -= 2
	fmt.Println("hasil = ", b)
	c *= 2
	fmt.Println("hasil = ", c)
	d /= 2
	fmt.Println("hasil = ", d)

	b++
	fmt.Println("hasil = ", b)
	c--
	fmt.Println("hasil = ", c)

}
package main

import "fmt"

func sum(numbers ...int) int{
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}

func main(){
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15))
	fmt.Println(sum(1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20))

	// jika dari slice 
	numbers := []int{1,2,3,4,5}
	fmt.Println(sum(numbers...))
}
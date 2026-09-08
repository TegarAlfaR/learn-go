package main

import "fmt"

func main(){
	counter := 1

	for counter <= 10{
		fmt.Println("perulangan ke -",counter)
		counter++
	}
	fmt.Println("selesai")

	for count := 1; count <= 5; count++{
		fmt.Println("CIHUY -",count)
	}

	// manual
	names := []string{"tegar", "alfa", "rizzi"}
	for i := 0; i < len(names); i++{
		fmt.Println("helo", names[i])
	}

	// for range
	for index, name := range names{
		fmt.Println("index", index, "=", name)
	}

	// kalo ngga mau dipakein key/indexnya
	for _, name := range names{
		fmt.Println(name)
	}
}
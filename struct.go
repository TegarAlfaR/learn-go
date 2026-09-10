package main

import "fmt"


// buat struct untuk field nya pake pascal case (Name, Age, dst)

type Customer struct{
	Name, Address string
	Age int
}

func main() {

	// cara ke 1
	var tegar Customer
	
	tegar.Name = "Tegar Alfa Rizzi"
	tegar.Address = "Bogor"
	tegar.Age = 22



	fmt.Println(tegar)
	fmt.Println(tegar.Name)
	fmt.Println(tegar.Address)
	fmt.Println(tegar.Age)

	// cara ke dua

	budi := Customer{
		Name: "Budi Junaidi",
		Address: "Jakarta",
		Age: 30,
	}

	fmt.Println(budi)
	fmt.Println(budi.Name)
	fmt.Println(budi.Address)
	fmt.Println(budi.Age)

	// cara ke 3
	joko := Customer{"Joko Jimy", "Bandung", 22}
	fmt.Println(joko)
	fmt.Println(joko.Name)
	fmt.Println(joko.Address)
	fmt.Println(joko.Age)
}
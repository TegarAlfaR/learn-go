package main

import "fmt"

func main() {

	// buat slice dari array
	names := [...]string{"Tegar", "Alfa", "Rizzi", "Andi", "Lunar", "Budi", "Utomo", "Bajigur"}

	slice1 := names[:]
	fmt.Println("slice semua : " ,slice1)

	slice2 := names[:4]
	fmt.Println("slice dari awal sampe index ke 4", slice2)

	slice3 := names[2:]
	fmt.Println("slice dari index ke 2 sampe akhir", slice3)

	slice4 := names[2:5]
	fmt.Println("slice dari index ke 2 sampe index ke 4  ", slice4)



	slice5 := []string{"jambu", "mangga", "apel", "pisang"}
	fmt.Println("slice langsung", slice5)


	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jum'at", "Sabtu", "Minggu"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	
	// akan ngerubah value/isi dari array days
	daySlice1[1] = "Minggu Cihuy"

	fmt.Println(daySlice1)
	fmt.Println(days)


	// jika arraynya sudah tidak bisa ditambah, maka akan ada array yg dibuat auto sama si slicenya (sayangnya array awal ngga keubah)
	daySlice2 := append(daySlice1, "Libur Cihuy")
	daySlice2[0] = "Sabtu Cihuy"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)


		// buat slice langsung
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Tegar"
	newSlice[1] = "Alfa"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	
	newSlice2 := append(newSlice, "Rizzi")
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "cihuy"

	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy daata slice

	fromSlice := days[:]

	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// kalo ingin langsung buat slice, tanda [] jangan diisi
	iniArray1 := [3]int{1,2,3}
	iniArray2 := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray1)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)
}
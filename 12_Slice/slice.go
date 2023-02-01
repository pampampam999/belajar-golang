package main

import "fmt"

func main() {

	var iniArray = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
	}

	slice1 := iniArray[2:4]

	fmt.Println(iniArray)
	fmt.Println(slice1)

	newSlice := make([]string, 2, 5) //deklarasi slice {tipe data, len , capasity}

	newSlice[0] = "Coba"
	newSlice[1] = "Saja"

	fmt.Println(newSlice)

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) //mengcopy newslice ke dalam variable copyslice
	fmt.Println(copySlice)

}

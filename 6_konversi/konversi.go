package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	//jika dari int tinggi di konversi ke int rendah maka jika melebihi batas atas akan terjadi over
	// dan melanjutkan dari bawah

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Billit"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}

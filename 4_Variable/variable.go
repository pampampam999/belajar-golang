package main

import "fmt"

func main() {
	var kata string = "ini varible string"
	var coba = 123

	// Variable tidak bisa di deklarasikan ulang
	//var kata = ""

	//variable dapat di perbarui datanya
	kata = "Yeay ganti"

	fmt.Println(kata)
	fmt.Println(coba)

}

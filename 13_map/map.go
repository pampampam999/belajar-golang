package main

import "fmt"

func main() {

	//deklarasi singkat
	person := map[string]string{
		"name":    "Billit",
		"address": "Surabaya",
	}

	//cara menmbahkan atau replace setelah deklarasi
	person["tittle"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["name"] = "Belajar Go"
	book["page"] = "200"
	book["coba"] = "apa aja deh"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "coba")

	fmt.Println(book)
	fmt.Println(len(book))
}

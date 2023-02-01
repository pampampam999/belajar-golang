package main

import "fmt"

func main() {

	pintu := "delapan"

	switch pintu {
	case "satu":
		fmt.Println("Pintu no satu")
	case "dua":
		fmt.Println("Pintu no Dua")
	default:
		fmt.Println("Ga ada pintunya")
	}

	//switch dengan short statment
	switch length := len(pintu); length > 5 {
	case true:
		fmt.Println("Terlalu banyak")
	case false:
		fmt.Println("Dikit")
	}
}

package main

import "fmt"

func main() {

	var name string = "pam"

	if name == "pam" {
		fmt.Println("Benar pam")
	} else if name == "oxy" {
		fmt.Println("ini oxy")
	} else {
		fmt.Println("Kamu siapa?")
	}

	//memberikan short statement sebelum kondisi dengan ;
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	}

}

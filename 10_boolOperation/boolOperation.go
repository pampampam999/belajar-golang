package main

import "fmt"

func main() {
	var nilaiUjian = 80
	var nilaiAbsen = 75

	var lulusUjian = nilaiUjian >= 80
	var lulusAbsen = nilaiAbsen >= 80

	var lulus = lulusUjian && lulusAbsen

	fmt.Println(lulus)

}

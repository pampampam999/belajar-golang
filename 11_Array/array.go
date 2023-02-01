package main

import "fmt"

func main() {

	var bulan [12]string = [12]string{
		"Januari", "Februari", "Maret",
		"April", "Mei", "Juni", "Juli",
		"Agustus", "September", "Oktober",
		"November", "Desember"}
	hari := [12]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	fmt.Println(bulan)
	fmt.Println(len(bulan))
	fmt.Println(hari)
	fmt.Println(bulan[1])
	bulan[1] = "Maret"
	fmt.Println(bulan[1])

}

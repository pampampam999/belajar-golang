package main

import "fmt"

func main() {
	type NoKTP string
	type Maried bool

	var ValueNoKTP NoKTP = "109238091928309123"
	var MariedStatus Maried = true

	fmt.Println(ValueNoKTP)
	fmt.Println(MariedStatus)
}

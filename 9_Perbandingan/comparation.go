package main

import "fmt"

func main() {
	var name1 = "Bil"
	var name2 = "Bil"

	var result = name1 == name2

	fmt.Println(result)

	fmt.Println(name1 > name2)
	fmt.Println(name1 < name2)
	fmt.Println(name1 >= name2)
	fmt.Println(name1 <= name2)
	fmt.Println(name1 != name2)
}

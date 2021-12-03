package main

import "fmt"

func main() {
	// Const Tidak Dapat Melakukan Re-Deklarasi Value
	// Cara Deklarasi Pertama
	const firstName string = "Dimas"

	// Cara Deklarasi Kedua
	const lastName = "Eka Adinandra"

	// Deklarasi Multiple Constant
	const (
		first string = "Dimas"
		last         = "Eka Adinandra"
	)

	fmt.Println(firstName, lastName)
	fmt.Println(first, last)
}

package main

import "fmt"

func main() {
	// Membuat Alias Untuk Tipe Data

	type Name string
	type Number int
	type Boolean bool

	var fullName Name = "Dimas Eka Adinandra"
	fmt.Println(fullName)
}

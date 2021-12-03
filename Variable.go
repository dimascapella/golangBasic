package main

import "fmt"

func main() {
	// Variable Dapat Di Passing Data Lagi Setelah Melakukan Deklarasi
	// Variable Dengan Tipe Data String
	var name string

	// Variable Dengan Tipe Data Number
	var age int

	// Variable Dengan Tipe Data Boolean
	var student bool

	// Cara Deklarasi Variable ke 2
	// var name = "Dimas Eka Adinandra"

	// Cara Deklarasi Variable k 3
	// name := "Dimas Eka Adinandra"

	// Cara Deklarasi Dengan Multiple Variable
	var (
		firstName = "Dimas"
		lastName  = "Eka Adinandra"
	)

	fmt.Println(firstName, lastName)

	name = "Dimas Eka Adinandra"
	age = 22
	student = false
	fmt.Print("My Name is ", name, " I'm ", age, " Old", " and My Status is ", student)
}

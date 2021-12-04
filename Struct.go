package main

import "fmt"

func main() {
	// First Option
	var data Biodata
	data.name = "Dimas Eka Adinandra"
	data.address = "Gunung Jati"
	data.age = 22
	fmt.Println(data)

	// Second Option
	personalData := Biodata{"Dimas", "Malang", 22}
	fmt.Println(personalData)

	// Third Option
	personal := Biodata{
		name:    "Eka",
		address: "Gunung Jati",
		age:     22,
	}
	fmt.Println(personal)
	personal.Hello("tara")
}

// Struct Merupakan template data seperti Database / Class
type Biodata struct {
	name, address string
	age           int
}

// Struct Function
func (biodata Biodata) Hello(name string) {
	println("Hallo,", name, " My Name is", biodata.name)
}

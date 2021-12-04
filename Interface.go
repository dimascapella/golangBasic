package main

import (
	"errors"
	"fmt"
)

// Interface merupakan tipe data abstract dan berisikan definisi method
type Person struct {
	Name, address string
}

type HasData interface {
	GetName() string
	GetAddress() string
}

func sayHay(hasData HasData) {
	println("My Name is", hasData.GetName(), "I'm Live at", hasData.GetAddress())
}

func (person Person) GetName() string {
	return person.Name
}

func (person Person) GetAddress() string {
	return person.address
}

// Interface Kosong
func emptyInterface(number int) interface{} {
	if number == 1 {
		return 1
	} else if number == 2 {
		return "Interface Kosong"
	} else {
		return true
	}
}

// Error Interface (Helper From Golang)
func divide(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi Tidak Boleh 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	biodata := Person{"Dimas", "Gunung Jati"}
	sayHay(biodata)

	// Interface kosong
	// Full Line (var number interface{} = emptyInterface(1))
	number := emptyInterface(1)
	fmt.Println(number)

	// Option 1
	fmt.Println(divide(1, 0))

	// Option 2
	hasil, err := divide(100, 10)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}
}

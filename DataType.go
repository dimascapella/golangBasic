package main

import "fmt"

func main() {
	// Data Type String
	stringValue := "Dimas"

	// Data Type int
	intValue := 10

	// Data Type Boolean
	boolValue := true
	fmt.Print(stringValue, intValue, boolValue)

	// Data Type Array
	var names = [3]string{"Dimas", "Eka", "Adinandra"}
	fmt.Println(len(names))

	// Data Type Slice
	var number = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := 1
	end := 8
	fmt.Println(number[start:end])
	fmt.Println(number[start:])
	fmt.Println(number[:end])
	fmt.Println(number[:])

	// function Slice
	// Len (get Length of Slice)
	fmt.Println(len(number[start:end]))

	// Cap (get Length of Capacity)
	fmt.Println(cap(number[start:end]))

	// Append (Insert New Data)
	number2 := append(number[:], 11)
	fmt.Println(number2)

	// Make (Create new Array with condition (Type of Array, Length of Array, Capacity of Array))
	newNumber := make([]int, 5, 5)
	newNumber[2] = 1000
	fmt.Println(newNumber)

	// Copy (Create a Copy of Array (To Variable, From Variable))
	initVar := make([]int, len(number2), cap(number2))
	copy(initVar, number2)
	fmt.Println(initVar)

	// Array [...] or [10] (with number)
	// Slice [] (Blank)

	// Data Type Map
	mapValue := map[int]string{
		0: "Dimas",
		1: "Eka",
		2: "Adinandra",
	}

	fmt.Println(mapValue[0])

	// Map Function
	// Make
	bioValue := map[string]string{
		"name":   "Dimas Eka Adinandra",
		"task":   "Learning Golang",
		"status": "Learning",
	}

	delete(bioValue, "status")
	fmt.Println(bioValue)
}

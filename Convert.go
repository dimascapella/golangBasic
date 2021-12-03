package main

import "fmt"

func main() {
	// Convert Int to Another Int
	nilai32 := 128
	nilai64 := int64(nilai32)
	nilai8 := int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	// Convert Int to String
	nilai := 69
	convert := string(nilai)

	fmt.Println(convert)

	// Convert Byte to String
	text := "Hallo"
	char := text[1]
	convertChar := string(char)

	fmt.Println(convertChar)
}

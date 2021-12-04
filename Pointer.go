package main

import (
	"fmt"
)

type Address struct {
	alamat, provinsi, negara string
}

func main() {
	// Pass By Value = Passing Semua Data Dan Memerlukan Memori yang sama ketika melakukan passing
	alt1 := Address{"Gunung Jati", "Jawa Timur", "Indonesia"}
	alt2 := alt1

	alt2.negara = "Germany"
	fmt.Println(alt1)
	fmt.Println(alt2)

	// Pass By Reference (Pointer) = Mereferensikan variable ke value/data yang sudah ada
	temp1 := Address{"Dampit", "Jawa Timur", "Indonesia"}
	// Baris Kode Asli
	// var temp2 *Address = &temp1
	temp2 := &temp1

	temp2.negara = "Australia"

	// Operator * (Bintang) melakukan pemindahan refer data
	// Case Pertama Pembuatan Data Baru
	// temp2 := &Address{"Malang", "Jatim", "Malaysia"}

	// Case Kedua Pemindahan semua refer ke dalam data baru
	*temp2 = Address{"Batu", "Jawa Timur", "Indonesia"}

	fmt.Println(temp1)
	fmt.Println(temp2)
}

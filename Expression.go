package main

func main() {
	// if else expression
	number := 5
	if number == 2 {
		println("Cocok")
	} else if number == 5 {
		println("Mantul")
	} else {
		println("Tidak Cocok")
	}

	// short Statement if else expression
	arr := []int{1, 2, 3}
	if i := len(arr); i >= 5 {
		println("Kebesaran")
	} else {
		println("Sesuai")
	}

	// Switch Expression
	name := "Eka Adinandra"
	switch name {
	case "Dimas":
		println("Sesuai")
	case "Eka Adinandra":
		println("Tidak Sesuai")
	default:
		println("Cek Kembali")
	}

	// Short Statement Switch Expression
	switch a := len(arr); a >= 3 {
	case true:
		println("Sesuai")
	case false:
		println("Tidak Sesuai")
	}

	// Switch Expression Without Condition
	length := len(arr)
	switch {
	case length > 5:
		println("Terlalu Panjang")
	case length <= 5:
		println("Sesuai")
	default:
		println("Format Salah")
	}

	// For Loop Expression
	// For Loop Same As While Loop ( Optional Condition "true" and "false")
	for true {
		println("Hallo")
	}

	// For Loop With Condition
	i := 1
	for i < 10 {
		println("Hallo ", i)
		i++
	}

	// For Loop With Statement (Java)
	for j := 1; j <= 5; j++ {
		println("Hallo ke ", j)
	}

	// For Loop With Statement (Python)
	for index, array := range arr {
		println(index, array)
	}

	// Break Condition
	start := 1
	for start <= 10 {
		if start == 8 {
			break
		}
		println(start)
		start += 1
	}

	// Continue Condition
	for q := 0; q <= 10; q++ {
		if q%2 == 0 {
			continue
		}
		println("Continue = ", q)
	}
}

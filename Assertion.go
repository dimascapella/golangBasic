package main

func random() interface{} {
	return "Holla"
}

func main() {
	result := random()
	resultString := result.(string)
	println(resultString)

	// resultInt := result.(int)
	// println(resultInt)

	// Switch Case Assertion
	switch temp := result.(type) {
	case string:
		println(temp, "is String")
	case int:
		println(temp, "is Int")
	default:
		println("Unknown")
	}
}

package main

import (
	"strings"
)

func main() {
	info()

	sumValue(20, 50)

	println(isPalindrome("Katak"))

	sum, times, minus := multiReturn(5, 2)
	println(sum, times, minus)

	divide, module := multiNamed(10, 2)
	println(divide, module)

	data := []int{10, 5, 2, 3, 5}
	println(timesAll(2, data...))

	Greeting := Hello
	println(Greeting("Dimas"))

	TextFiltering("Anjing", FilterFunction)
}

// Function Without Parameter
func info() {
	var (
		name    = "Dimas Eka Adinandra"
		age     = 22
		address = "Gunung Jati"
	)

	println("Hello My Name is ", name, " I'm ", age, " Old and I Live at ", address)
}

// Function With Parameter
func sumValue(val1 int, val2 int) {
	println("Result : ", val1+val2)
}

// Function With Return
func isPalindrome(text string) bool {
	text = strings.ToLower(text)
	reversedText := reverse(text)
	palin := true
	if text != reversedText {
		palin = false
	}
	return palin
}

func reverse(text string) string {
	var result string
	for _, v := range text {
		result = string(v) + result
	}
	return result
}

// Function With Multiple Return Value
func multiReturn(val1 int, val2 int) (int, int, int) {
	sum := val1 + val2
	times := val1 * val2
	minus := val1 - val2
	return sum, times, minus
}

// Function Return Named Values
func multiNamed(val1 int, val2 int) (divide, module int) {
	divide = val1 / val2
	module = val1 % val2
	return divide, module
}

// Variadic Function (a Paramater With Multiple Value (array))
func timesAll(a int, b ...int) int {
	result := 0
	for _, value := range b {
		result += (a * value)
	}
	return result
}

// Function as Value
func Hello(text string) string {
	return "Hello " + text
}

// Function as Parameter
type Filter func(string) string

func FilterFunction(text string) string {
	text = strings.ToLower(text)
	if text == "anjing" {
		return "..."
	} else {
		return text
	}
}

func TextFiltering(text string, filterFunc Filter) {
	text = filterFunc(text)
	println("Filtered Text :", text)
}

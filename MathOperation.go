package main

import "fmt"

func main() {
	sum := 10
	sum += 10
	fmt.Println(sum)

	times := 5
	times *= 2
	fmt.Println(times)

	divide := 10
	divide /= 2
	fmt.Println(divide)

	minus := 10
	minus -= 5
	fmt.Println(minus)
}

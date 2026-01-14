package main

import "fmt"

// tipe data constant tidak boleh diubah nilainya
func main() {
	const name string = "Golang"
	const lastName = "Programming"

	fmt.Println(name)
	fmt.Println(lastName)

	// constant dengan multiple value
	const (
		firstName  = "Reyhand"
		middleName = "Astra"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
}

package main

import "fmt"

func main() {
	// var name string

	// var name = "Reyhand Astra"
	name := "Reyhand Astra" //cara singkat deklarasi variable
	fmt.Println(name)

	name = "Resa"
	fmt.Println(name)

	//menggunakan multiple variable
	var (
		firstName = "Reyhand"
		lastName  = "Astra"
	)

	fmt.Println(firstName + " " + lastName)
	// fmt.Println(lastName)
}

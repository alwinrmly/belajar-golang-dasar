package main

import "fmt"

func main() {
	const (
		firstName string = "Alwin"
		lastName         = "Sasmita"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Asep"
	// lastName = "Wicaksono"
}

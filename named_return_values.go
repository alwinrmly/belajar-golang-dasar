package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Alwin"
	middleName = "Ramli"
	lastName = "Sasmita"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}

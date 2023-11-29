package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	// person["name"] = "Alwin"
	// person["address"] = "Ciledug"

	person := map[string]string{
		"name":    "Alwin",
		"address": "ciledug",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Alwin"
	book["Ups"] = "Wrong"

	fmt.Println(book)

	delete(book, "Ups")

	fmt.Println(book)
}

package main

import "fmt"

func main() {
	name := "Daryono"

	if name == "Alwin" {
		fmt.Println("Hello Alwin")
	} else if name == "Tejo" {
		fmt.Println("Hello Tejo")
	} else if name == "Agung" {
		fmt.Println("Hello Agung")
	} else {
		fmt.Println("Hi, Perkenalkan Nama Anda?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")

	}
}

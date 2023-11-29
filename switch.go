package main

import "fmt"

func main() {
	name := "Alwin"

	switch name {
	case "Alwin":
		fmt.Println("Hello Alwin")
	case "Tejo":
		fmt.Println("Hello Tejo")
	default:
		fmt.Println("Hi, Perkenalkan Nama Anda?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	name = "DarsonoDarsonoDarsono"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}

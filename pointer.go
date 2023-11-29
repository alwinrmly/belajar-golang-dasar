package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Ciledug", "Banten", "Indonesia"}
	// address2 := address1 // copy value

	// address2.City = "Depok"
	// address2.Province = "Jawa Barat"

	// fmt.Println(address1) // tidak berubah
	// fmt.Println(address2) // City berubah menjadi Depok dan Province berubah menjadi Jawa Barat

	var address1 Address = Address{"Ciledug", "Banten", "Indonesia"}
	var address2 *Address = &address1 // reference value

	address2.City = "Depok"
	address2.Province = "Jawa Barat"

	fmt.Println(address1) // data akan berubah
	fmt.Println(address2) // City berubah menjadi Depok dan Province berubah menjadi Jawa Barat
}

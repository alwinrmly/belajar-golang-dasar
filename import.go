package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Alwin")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak dapat diakses
	// fmt.Println(helper.sayGoodBye) // tidak dapat diakses
}

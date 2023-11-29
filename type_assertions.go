package main

import (
	"fmt"
)

func random() any {
	return false
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// var resultInt int result.(int) // akan panic
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}

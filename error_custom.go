package main

import (
	"fmt"
)

type validationError struct {
	message string
}

func (v *validationError) Error() string {
	return v.message
}

type notFoundError struct {
	message string
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "Alwin" {
		return &notFoundError{"Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("Paijo", nil)
	if err != nil {
		// terjadi error
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("Unknown Error:", finalError.Error())
		}

	} else {
		// Sukses
		fmt.Println("Sukses")
	}
}

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Error!")
	}
}

func main() {
	runApp(true)
	fmt.Println("Alwin Ramli Sasmita")
}

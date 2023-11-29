package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Alwin")
	fmt.Println(result)

	fmt.Println(getHello("Tejo"))
	fmt.Println(getHello("Asep"))
}

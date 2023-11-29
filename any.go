package main

import "fmt"

func Ups() any {
	// return 1
	// return 31
	return "Banteng"
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}

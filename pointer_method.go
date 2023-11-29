package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	alwin := Man{"Alwin"}
	alwin.Married()

	fmt.Println(alwin.Name)
}

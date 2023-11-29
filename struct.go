package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var alwin Customer
	alwin.Name = "Alwin Ramli"
	alwin.Address = "Tangerang"
	alwin.Age = 25

	fmt.Println(alwin)
	fmt.Println(alwin.Name)
	fmt.Println(alwin.Address)
	fmt.Println(alwin.Age)

	tejo := Customer{
		Name:    "Tejo",
		Address: "Tangerang",
		Age:     40,
	}
	fmt.Println(tejo)

	asep := Customer{"Asep", "Tangerang", 31}
	fmt.Println(asep)

	alwin.sayHello("Dodit")
	tejo.sayHello("Dodit")
	asep.sayHello("Dodit")

}

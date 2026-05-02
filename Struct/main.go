package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var alex person // declaring a variable of type person
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	alex = person{firstName: "Alex", lastName: "Anderson", age: 30}
	alex.firstName = "Alexander" // updating the firstName field of the alex struct
	fmt.Println(alex)            // accessing the firstName field of the alex struct and printing it
}

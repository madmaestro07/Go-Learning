package main

import "fmt"

func test() {
	fmt.Println("Hi there, welcome to Go!")
}

func main() {
	var name string = "My name is maestro"
	// name := "My name is maestro" // short declaration of new variable, only works inside functions
	name = "My name is not maestro" // reassigning a new value to the variable
	fmt.Println(name)

	// declaring multiple variables at once and assigning values to them
	names := list{"maestro", "john", "doe", value_add()} // declaring a slice of type list and initializing it with some values
	names = append(names, value_add()) // appending a new value to the slice
	names.print() //  names is of type list and we can call the print method on it to print the values in the list

}

func value_add() string{
	return  "oliver"
}
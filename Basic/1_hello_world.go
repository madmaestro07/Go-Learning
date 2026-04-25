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
	names := []string{"maestro", "john", "doe", value_add()}
	names = append(names, value_add()) // appending a new value to the slice

	for i,name := range names {
		fmt.Println(i, name)
	}

}

func value_add() string{
	return  "oliver"
}
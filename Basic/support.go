package main

type list []string // defining a new type called list which is a slice of strings - we can call methods on this type to perform operations on the list of strings

func (l list) print() {  // defining a method called print on the list type which takes a list as a receiver and prints the values in the list
	for i, name := range l {
		println(i, name)
	}
}
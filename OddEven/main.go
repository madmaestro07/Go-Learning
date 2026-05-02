package main

import "fmt"

func main() {
	numSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // creating a slice of integers from 0 to 10

	for _, num := range numSlice { // iterating over the slice of integers
		if num%2 == 0 && num != 0 { // checking if the number is even
			fmt.Printf("%v is even\n", num) // printing that the number is even
		} else if num%2 != 0 { // checking if the number is odd
			fmt.Printf("%v is odd\n", num) // printing that the number is odd
		} else { // if the number is 0
			fmt.Printf("%v is neither even nor odd\n", num) // printing that the number is neither even nor odd
		}
	}
}

package main

import "fmt"

func main() {
	var feet float64

	fmt.Printf("Enter the distance/length in Feet: ")
	fmt.Scanf("%f", &feet)

	var output = feet * 0.3408

	fmt.Printf("In meters it is %f\n", output)

}

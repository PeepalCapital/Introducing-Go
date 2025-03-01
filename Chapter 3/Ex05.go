package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Printf("Enter the temperature in Fahrenheit: ")
	fmt.Scanf("%f", &fahrenheit)

	var output = (fahrenheit - 32) * (5.0 / 9.0)

	fmt.Printf("In Celsius it is %f\n", output)

}

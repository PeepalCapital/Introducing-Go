package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min_value := x[0]

	for _, value := range x {
		if value < min_value {
			min_value = value
		}
	}
	fmt.Println("Min Value is: ", min_value)

}

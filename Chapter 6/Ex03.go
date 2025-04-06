/* There is no inbuilt max function in Golang */

package main

import "fmt"

func greatest(args ...int) int {
	maxval := args[0]
	for _, num := range args {
		if num > maxval {
			maxval = num
		}
	}
	return maxval
}

func main() {
	fmt.Println(greatest(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

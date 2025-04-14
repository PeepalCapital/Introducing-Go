/* A function to swap 2 integers */

package main

import "fmt"

func swap(x *int, y *int) (int, int) {
	return *y, *x
}

func main() {
	a := 2
	b := 3
	fmt.Println(swap(&a, &b))
}

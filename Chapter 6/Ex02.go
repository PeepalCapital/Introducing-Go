package main

import "fmt"

func halfeven(x int) (int, bool) {
	if x%2 == 0 {
		return x / 2, true
	} else {
		return x / 2, false
	}

}

func main() {
	fmt.Println(halfeven(2))
}

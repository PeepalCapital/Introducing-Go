package main

import "fmt"

func main() {
	fmt.Println("(true && false) || (false && true) || ! (false && false) is ", (true && false) || (false && true) || !(false && false))
}

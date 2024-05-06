package main

import "fmt"

func main() {
	fmt.Println(sum(10, 10))
	fmt.Println(multiply(10, 10))
}

func sum(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

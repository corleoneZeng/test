package main

import "fmt"

func main() {
	fmt.Println("Git is a distributed version control system")
	fmt.Println("Git is free software distributed under the GPL")
	sum := add(1, 2)
	fmt.Println("sum=", sum)
}

func add(a, b int) int {
	return a + b
}

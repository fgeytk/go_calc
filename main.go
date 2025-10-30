package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(add(5,3))
	fmt.Println(subtract(10,4))
	fmt.Println(multiply(3,7))
	fmt.Println(divide(20,4))
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) (int, error) {
	if b == 0 {
	return 0, fmt.Errorf("division par zero")
	}	else {
		return a / b, nil
	}}

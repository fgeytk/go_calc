package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(add(5, 3))
	fmt.Println(subtract(10, 4))
	fmt.Println(multiply(3, 7))
	fmt.Println(divide(20, 4))

	// Choix de l'operation

	fmt.Println("Deux nombres")
	var a, b int
	fmt.Print("Entrez le premier nombre: ")
	fmt.Scanln(&a)
	fmt.Print("Entrez le deuxieme nombre: ")
	fmt.Scanln(&b)
	fmt.Printf("Choix de l'operation")
	fmt.Println("1. Addition")
	fmt.Println("2. Soustraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	var choice int
	fmt.Print("Entrez votre choix (1-4): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Printf("Resultat: %d\n", add(a, b))
	case 2:
		fmt.Printf("Resultat: %d\n", subtract(a, b))
	case 3:
		fmt.Printf("Resultat: %d\n", multiply(a, b))
	case 4:
		result, err := divide(a, b)
		if err != nil {
			fmt.Println("Erreur:", err)
		} else {
			fmt.Printf("Resultat: %d\n", result)
		}
	}
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
	} else {
		return a / b, nil
	}
}

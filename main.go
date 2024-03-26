package main

import (
	"fmt"
	"os"
)

func main() {
	// Verifies if there were parameters passed
	if len(os.Args) < 2 {
		fmt.Println("Use: Hello <nombre>")
		os.Exit(1)
	}

	// OObtains the name passed as an argument
	nombre := os.Args[1]

	// Prints the greeting
	fmt.Printf("Hola, %s!\n", nombre)
}

package main

import "fmt"

func forgeron(c *Character) {
	for {
		fmt.Println("\n--- Forgeron ---")
		fmt.Println("1. Chapeau de l’aventurier")
		fmt.Println("2. Tunique de l’aventurier")
		fmt.Println("3. Bottes de l’aventurier")
		fmt.Println("4. Retour")

		var choix int
		fmt.Scan(&choix)

		switch choix {
		case 4:
			return
		default:
			fmt.Println(" ")
		}
	}
}
